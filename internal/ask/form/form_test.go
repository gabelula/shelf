package form_test

import (
	"testing"
	"time"

	"github.com/ardanlabs/kit/tests"
	"github.com/coralproject/shelf/internal/ask/form"
	"github.com/coralproject/shelf/internal/ask/form/formfix"
	"github.com/coralproject/shelf/internal/platform/db"
)

// prefix is what we are looking to delete after the test.
const prefix = "FTEST"

func setupForms(t *testing.T, fixture string) ([]form.Form, *db.DB) {
	tests.ResetLog()

	fms, err := formfix.Get("form")
	if err != nil {
		t.Fatalf("%s\tShould be able retrieve form fixture : %s", tests.Failed, err)
	}
	t.Logf("%s\tShould be able retrieve form fixture.", tests.Success)

	db, err := db.NewMGO(tests.Context, tests.TestSession)
	if err != nil {
		t.Fatalf("Should be able to get a Mongo session : %v", err)
	}

	return fms, db
}

func teardownForms(t *testing.T, db *db.DB) {
	if err := formfix.Remove(tests.Context, db, prefix); err != nil {
		t.Fatalf("%s\tShould be able to remove the forms : %v", tests.Failed, err)
	}
	t.Logf("%s\tShould be able to remove the forms.", tests.Success)

	db.CloseMGO(tests.Context)
	tests.DisplayLog()
}

func TestUpsertDelete(t *testing.T) {
	fms, db := setupForms(t, "form")
	defer teardownForms(t, db)

	t.Log("Given the need to upsert and delete forms.")
	{
		t.Log("\tWhen starting from an empty forms collection")
		{
			//----------------------------------------------------------------------
			// Upsert the form.

			if err := form.Upsert(tests.Context, db, &fms[0]); err != nil {
				t.Fatalf("\t%s\tShould be able to upsert a form : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to upsert a form.", tests.Success)

			//----------------------------------------------------------------------
			// Get the form.

			fm, err := form.Retrieve(tests.Context, db, fms[0].ID.Hex())
			if err != nil {
				t.Fatalf("\t%s\tShould be able to get the form by id : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to get the form by id.", tests.Success)

			//----------------------------------------------------------------------
			// Check that we got the form we expected.

			if fms[0].ID.Hex() != fm.ID.Hex() {
				t.Fatalf("\t%s\tShould be able to get back the same form.", tests.Failed)
			}
			t.Logf("\t%s\tShould be able to get back the same form.", tests.Success)

			//----------------------------------------------------------------------
			// Delete the form.

			if err := form.Delete(tests.Context, db, fms[0].ID.Hex()); err != nil {
				t.Fatalf("\t%s\tShould be able to delete the form : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to delete the form.", tests.Success)

			//----------------------------------------------------------------------
			// Get the form.

			_, err = form.Retrieve(tests.Context, db, fms[0].ID.Hex())
			if err == nil {
				t.Fatalf("\t%s\tShould generate an error when getting a form with the deleted id : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould generate an error when getting a form with the deleted id.", tests.Success)

			//----------------------------------------------------------------------
			// Create a new fresh form.

			fms[0].ID = ""

			if err := form.Upsert(tests.Context, db, &fms[0]); err != nil {
				t.Fatalf("\t%s\tShould be able to upsert a form : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to upsert a form.", tests.Success)

			//----------------------------------------------------------------------
			// Ensure that an ID was set.

			if fms[0].ID == "" {
				t.Fatalf("\t%s\tShould be able to add an ID when upserting a new form : ID was not assigned", tests.Failed)
			}
			t.Logf("\t%s\tShould be able to add an ID when upserting a new form.", tests.Success)

			//----------------------------------------------------------------------
			// Get the form.

			fm, err = form.Retrieve(tests.Context, db, fms[0].ID.Hex())
			if err != nil {
				t.Fatalf("\t%s\tShould be able to get the form by id : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to get the form by id.", tests.Success)

			//----------------------------------------------------------------------
			// Check that we got the form we expected.

			if fms[0].ID.Hex() != fm.ID.Hex() {
				t.Fatalf("\t%s\tShould be able to get back the same form.", tests.Failed)
			}
			t.Logf("\t%s\tShould be able to get back the same form.", tests.Success)

			//----------------------------------------------------------------------
			// Delete the form.

			if err := form.Delete(tests.Context, db, fms[0].ID.Hex()); err != nil {
				t.Fatalf("\t%s\tShould be able to delete the form : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to delete the form.", tests.Success)

			//----------------------------------------------------------------------
			// Get the form.

			_, err = form.Retrieve(tests.Context, db, fms[0].ID.Hex())
			if err == nil {
				t.Fatalf("\t%s\tShould generate an error when getting a form with the deleted id : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould generate an error when getting a form with the deleted id.", tests.Success)
		}
	}
}

func TestList(t *testing.T) {
	fms, db := setupForms(t, "form")
	defer teardownForms(t, db)

	t.Log("Given the need to upsert and delete forms.")
	{
		t.Log("\tWhen starting from an empty forms collection")
		{

			//----------------------------------------------------------------------
			// Upsert the forms.

			if err := form.Upsert(tests.Context, db, &fms[0]); err != nil {
				t.Fatalf("\t%s\tShould be able to upsert a form : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to upsert forms.", tests.Success)

			// This particular looping logic is required at the moment due to issues
			// with the current CI environment.

			limit := 10
			offset := 0

			for {
				lfms, err := form.List(tests.Context, db, limit, offset)
				if err != nil {
					t.Fatalf("\t%s\tShould be able to list forms : %s", tests.Failed, err)
				}
				t.Logf("\t%s\tShould be able to list forms", tests.Success)

				if len(lfms) > limit {
					t.Fatalf("\t%s\tShould only return a maximum of the provided limit : Expected less than %d, got %d", tests.Failed, limit, len(lfms))
				}
				t.Logf("\t%s\tShould only return a maximum of the provided limit.", tests.Success)

				// If we are another page of data (we have itterated at least once) and
				// the length of the results is zero, then we have no more results to
				// paginate over. At this point, we know that we couldn't find all the
				// upserted records.
				if offset > 0 && len(lfms) == 0 {
					t.Fatalf("\t%s\tShould be able to find a form that was upserted : Could not an upserted form in result set", tests.Failed)
				}

				found := false
				for _, fm := range lfms {
					if fm.ID.Hex() == fms[0].ID.Hex() {
						found = true
						break
					}
				}

				// If not all the elements have been found yet, we need to continue the
				// itteration process.
				if !found {
					// Increase the offset by the limit amount.
					offset = offset + limit

					// Continue itteration.
					continue
				}

				// Exit the itteration in the event that all the upserted documents have
				// been found.
				break
			}
			t.Logf("\t%s\tShould be able to find a form that was upserted", tests.Success)
		}
	}
}

func TestUpdateStatus(t *testing.T) {
	fms, db := setupForms(t, "form")
	defer teardownForms(t, db)

	t.Log("Given the need to upsert and delete forms.")
	{
		t.Log("\tWhen starting from an empty forms collection")
		{

			//----------------------------------------------------------------------
			// Upsert the form.

			if err := form.Upsert(tests.Context, db, &fms[0]); err != nil {
				t.Fatalf("\t%s\tShould be able to upsert a form : %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould be able to upsert a form.", tests.Success)

			//----------------------------------------------------------------------
			// Update it's status.

			newStatus := "updated_" + time.Now().String()

			fm, err := form.UpdateStatus(tests.Context, db, fms[0].ID.Hex(), newStatus)
			if err != nil {
				t.Logf("\t%s\tShould be able to update a forms status without error : %s", tests.Success, err.Error())
			}
			t.Logf("\t%s\tShould be able to update a forms status without error.", tests.Success)

			//----------------------------------------------------------------------
			// Check we got the right form.

			if fm.ID.Hex() != fms[0].ID.Hex() {
				t.Fatalf("\t%s\tShould be able to retrieve a form given it's id : ID's of retrieved forms do not match", tests.Success)
			}
			t.Logf("\t%s\tShould be able to retrieve a form given it's id.", tests.Success)

			//----------------------------------------------------------------------
			// Check that it's status is changed.

			if fm.Status != newStatus {
				t.Fatalf("\t%s\tShould be able to set the status on the returned form : Expected %s, got %s", tests.Success, newStatus, fm.Status)
			}
			t.Logf("\t%s\tShould be able to set the status on the returned form.", tests.Success)

			//----------------------------------------------------------------------
			// Get a copy from the DB.

			rfm, err := form.Retrieve(tests.Context, db, fm.ID.Hex())
			if err != nil {
				t.Fatalf("\t%s\tShould be able to retrieve a form given it's id : %s", tests.Success, err.Error())
			}
			t.Logf("\t%s\tShould be able to retrieve a form given it's id.", tests.Success)

			//----------------------------------------------------------------------
			// Check that the DB copy has it's status changed.

			if rfm.Status != newStatus {
				t.Fatalf("\t%s\tShould be able to update a form's status in the database : Expected %s, got %s", tests.Failed, newStatus, rfm.Status)
			}
			t.Logf("\t%s\tShould be able to update a form's status in the database.", tests.Success)
		}
	}
}

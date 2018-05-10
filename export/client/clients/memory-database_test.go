//
// Copyright (c) 2018 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package clients

import (
	"testing"

	"github.com/edgexfoundry/edgex-go/export"
)

func testDB(t *testing.T, db DBClient) {
	r := export.Registration{}
	r.Name = "name"

	id, err := db.AddRegistration(&r)
	if err != nil {
		t.Fatalf("Error adding registration %v: %v", r, err)
	}

	regs, err := db.Registrations()
	if err != nil {
		t.Fatalf("Error getting registrations %v", err)
	}
	if len(regs) != 1 {
		t.Fatalf("There should be only one registration instead of %d", len(regs))
	}
	r2, err := db.RegistrationById(id.Hex())
	if err != nil {
		t.Fatalf("Error getting registrations by id %v", err)
	}
	if r2.ID.Hex() != id.Hex() {
		t.Fatalf("Id does not match %s - %s", r2.ID, id)
	}
	_, err = db.RegistrationById("INVALID")
	if err == nil {
		t.Fatalf("Registration should not be found")
	}

	r3, err := db.RegistrationByName(r.Name)
	if err != nil {
		t.Fatalf("Error getting registrations by name %v", err)
	}
	if r3.Name != r.Name {
		t.Fatalf("Id does not match %s - %s", r2.ID, id)
	}
	_, err = db.RegistrationByName("INVALID")
	if err == nil {
		t.Fatalf("Registration should not be found")
	}

	err = db.DeleteRegistrationById("INVALID")
	if err == nil {
		t.Fatalf("Registration should not be deleted")
	}

	err = db.DeleteRegistrationById(id.Hex())
	if err != nil {
		t.Fatalf("Registration should be deleted: %v", err)
	}

	id, err = db.AddRegistration(&r)
	if err != nil {
		t.Fatalf("Error adding registration %v: %v", r, err)
	}

	r.ID = id
	r.Name = "name2"
	err = db.UpdateRegistration(r)
	if err != nil {
		t.Fatalf("Error updating registration %v", err)
	}

	err = db.DeleteRegistrationByName("INVALID")
	if err == nil {
		t.Fatalf("Registration should not be deleted")
	}

	err = db.DeleteRegistrationByName(r.Name)
	if err != nil {
		t.Fatalf("Registration should be deleted: %v", err)
	}

	err = db.UpdateRegistration(r)
	if err == nil {
		t.Fatalf("Update should return error")
	}

}

func TestMemoryDB(t *testing.T) {
	memory := &memDB{}
	testDB(t, memory)
}
/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package lints

import (
	"testing"
)

func TestCAOrgNameBlank(t *testing.T) {
	inputPath := "../testlint/testCerts/caOrgNameEmpty.pem"
	expected := Error
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCAOrgNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caOrgNameMissing.pem"
	expected := Error
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCAOrgNameValid(t *testing.T) {
	inputPath := "../testlint/testCerts/caValOrgName.pem"
	expected := Pass
	out := Lints["e_ca_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

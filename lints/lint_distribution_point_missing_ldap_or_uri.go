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

/************************************************
RFC 5280: 4.2.1.13
When present, DistributionPointName SHOULD include at least one LDAP or HTTP URI.
************************************************/

package lints

import (
	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type distribNoLDAPorURI struct{}

func (l *distribNoLDAPorURI) Initialize() error {
	return nil
}

func (l *distribNoLDAPorURI) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *distribNoLDAPorURI) Execute(c *x509.Certificate) *LintResult {
	for _, point := range c.CRLDistributionPoints {
		if point = strings.ToLower(point); strings.HasPrefix(point, "http://") || strings.HasPrefix(point, "ldap://") {
			return &LintResult{Status: Pass}
		}
	}
	return &LintResult{Status: Warn}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_distribution_point_missing_ldap_or_uri",
		Description:   "When present in the CRLDistributionPoints extension, DistributionPointName SHOULD include at least one LDAP or HTTP URI",
		Citation:      "RFC 5280: 4.2.1.13",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &distribNoLDAPorURI{},
	})
}

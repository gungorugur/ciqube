package sonarqube

var metricDefinitionMap map[string]string

func init() {
	metricDefinitionMap = map[string]string{
		"complexity":                         "Complexity",
		"cognitive_complexity":               "Cognitive Complexity",
		"duplicated_blocks":                  "Duplicated blocks",
		"duplicated_files":                   "Duplicated files",
		"duplicated_lines":                   "Duplicated lines",
		"duplicated_lines_density":           "Duplicated lines (%)",
		"new_violations":                     "New issues",
		"new_xxx_violations":                 "New xxx issues",
		"violations":                         "Issues",
		"confirmed_issues":                   "Confirmed issues",
		"reopened_issues":                    "Reopened issues",
		"code_smells":                        "Code Smells",
		"new_code_smells":                    "New Code Smells",
		"sqale_rating":                       "Maintainability Rating",
		"sqale_index":                        "Technical Debt",
		"new_technical_debt":                 "Technical Debt on New Code",
		"sqale_debt_ratio":                   "Technical Debt Ratio",
		"new_sqale_debt_ratio":               "Technical Debt Ratio on New Code",
		"alert_status":                       "Quality Gate Status",
		"quality_gate_details":               "Quality Gate Details",
		"bugs":                               "Bugs",
		"new_bugs":                           "New Bugs",
		"reliability_rating":                 "Reliability Rating",
		"reliability_remediation_effort":     "Reliability remediation effort",
		"new_reliability_remediation_effort": "Reliability remediation effort on new code",
		"vulnerabilities":                    "Vulnerabilities",
		"new_vulnerabilities":                "Vulnerabilities on new code",
		"security_rating":                    "Security Rating",
		"security_remediation_effort":        "Security remediation effort",
		"new_security_remediation_effort":    "Security remediation effort on new code",
		"security_hotspots":                  "Security Hotspots",
		"new_security_hotspots":              "Security Hotspots on new code",
		"security_review_rating":             "Security Review Rating",
		"new_security_review_rating":         "Security Review Rating on new code",
		"security_hotspots_reviewed":         "Security Hotspots Reviewed",
		"classes":                            "Classes",
		"comment_lines":                      "Comment lines",
		"comment_lines_density":              "Comments (%)",
		"directories":                        "Directories",
		"files":                              "Files",
		"lines":                              "Lines",
		"ncloc":                              "Lines of code",
		"ncloc_language_distribution":        "Lines of code per language",
		"projects":                           "Projects",
		"statements":                         "Statements",
		"branch_coverage":                    "Condition coverage",
		"new_branch_coverage":                "Condition coverage on new code",
		"branch_coverage_hits_data":          "Condition coverage hits",
		"conditions_by_line":                 "Conditions by line",
		"covered_conditions_by_line":         "Covered conditions by line",
		"coverage":                           "Coverage",
		"new_coverage":                       "Coverage on new code",
		"line_coverage":                      "Line coverage",
		"new_line_coverage":                  "Line coverage on new code",
		"coverage_line_hits_data":            "Line coverage hits",
		"lines_to_cover":                     "Lines to cover",
		"new_lines_to_cover":                 "Lines to cover on new code",
		"skipped_tests":                      "Skipped unit tests",
		"uncovered_conditions":               "Uncovered conditions",
		"new_uncovered_conditions":           "Uncovered conditions on new code",
		"uncovered_lines":                    "Uncovered lines",
		"new_uncovered_lines":                "Uncovered lines on new code",
		"tests":                              "Unit tests",
		"test_execution_time":                "Unit tests duration",
		"test_errors":                        "Unit test errors",
		"test_failures":                      "Unit test failures",
		"test_success_density":               "Unit test success density (%)",
	}
}

//GetMetricDefinition returns metric definition from metric key
func GetMetricDefinition(metrickey string) string {

	definition := metricDefinitionMap[metrickey]
	if definition == "" {
		return metrickey
	}

	return definition
}

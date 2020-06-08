package wordcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCamelCase          = "camelCase"
	testDotCase            = "dot.case"
	testKebabCase          = "kebab-case"
	testPascalCase         = "PascalCase"
	testScreamingSnakeCase = "SCREAMING_SNAKE_CASE"
	testSnakeCase          = "snake_case"
	testTitleCase          = "Title Case"
	testWordCase           = "word case"
)

var testFunctions = []string{
	testCamelCase,
	testDotCase,
	testKebabCase,
	testPascalCase,
	testScreamingSnakeCase,
	testSnakeCase,
	testTitleCase,
	testWordCase,
}

var fnMap = map[string]func(string) string{
	testSnakeCase:          SnakeCase,
	testKebabCase:          KebabCase,
	testDotCase:            DotCase,
	testScreamingSnakeCase: ScreamingSnakeCase,
	testCamelCase:          CamelCase,
	testPascalCase:         PascalCase,
	testWordCase:           Words,
	testTitleCase:          TitleCase,
}

var testCases = map[string]map[string]string{
	"a": {
		testCamelCase:          "a",
		testPascalCase:         "A",
		testDotCase:            "a",
		testKebabCase:          "a",
		testScreamingSnakeCase: "A",
		testSnakeCase:          "a",
		testTitleCase:          "A",
		testWordCase:           "a",
	},
	"A": {
		testCamelCase:          "a",
		testPascalCase:         "A",
		testDotCase:            "a",
		testKebabCase:          "a",
		testScreamingSnakeCase: "A",
		testSnakeCase:          "a",
		testTitleCase:          "A",
		testWordCase:           "A",
	},
	"a---------b": {
		testCamelCase:          "aB",
		testPascalCase:         "AB",
		testDotCase:            "a.b",
		testKebabCase:          "a-b",
		testScreamingSnakeCase: "A_B",
		testSnakeCase:          "a_b",
		testTitleCase:          "A B",
		testWordCase:           "a b",
	},
	"*a.B": {
		testCamelCase:          "aB",
		testPascalCase:         "AB",
		testDotCase:            "a.b",
		testKebabCase:          "a-b",
		testScreamingSnakeCase: "A_B",
		testSnakeCase:          "a_b",
		testTitleCase:          "A B",
		testWordCase:           "a B",
	},
	"dooker": {
		testCamelCase:          "dooker",
		testPascalCase:         "Dooker",
		testDotCase:            "dooker",
		testKebabCase:          "dooker",
		testScreamingSnakeCase: "DOOKER",
		testSnakeCase:          "dooker",
		testTitleCase:          "Dooker",
		testWordCase:           "dooker",
	},
	"dookerSpam99_rawr": {
		testCamelCase:          "dookerSpam99Rawr",
		testScreamingSnakeCase: "DOOKER_SPAM99_RAWR",
		testDotCase:            "dooker.spam99.rawr",
		testKebabCase:          "dooker-spam99-rawr",
		testPascalCase:         "DookerSpam99Rawr",
		testSnakeCase:          "dooker_spam99_rawr",
		testTitleCase:          "Dooker Spam99 Rawr",
		testWordCase:           "dooker Spam99 rawr",
	},
	"IDOne_XMLHttp_ON": {
		testCamelCase:          "idOneXMLHTTPOn",
		testPascalCase:         "IdOneXmlHttpOn",
		testDotCase:            "id.one.xml.http.on",
		testKebabCase:          "id-one-xml-http-on",
		testScreamingSnakeCase: "ID_ONE_XML_HTTP_ON",
		testSnakeCase:          "id_one_xml_http_on",
		testTitleCase:          "ID One XML HTTP On",
		testWordCase:           "ID One XML Http ON",
	},
	"maxID": {
		testCamelCase:          "maxID",
		testPascalCase:         "MaxId",
		testDotCase:            "max.id",
		testKebabCase:          "max-id",
		testScreamingSnakeCase: "MAX_ID",
		testSnakeCase:          "max_id",
		testTitleCase:          "Max ID",
		testWordCase:           "max ID",
	},
	"maxId": {
		testCamelCase:          "maxID",
		testPascalCase:         "MaxId",
		testDotCase:            "max.id",
		testKebabCase:          "max-id",
		testScreamingSnakeCase: "MAX_ID",
		testSnakeCase:          "max_id",
		testTitleCase:          "Max ID",
		testWordCase:           "max Id",
	},
	"ENV_VAR": {
		testCamelCase:          "envVar",
		testPascalCase:         "EnvVar",
		testDotCase:            "env.var",
		testKebabCase:          "env-var",
		testScreamingSnakeCase: "ENV_VAR",
		testSnakeCase:          "env_var",
		testTitleCase:          "Env Var",
		testWordCase:           "ENV VAR",
	},
	"snake_case": {
		testCamelCase:          "snakeCase",
		testPascalCase:         "SnakeCase",
		testDotCase:            "snake.case",
		testKebabCase:          "snake-case",
		testScreamingSnakeCase: "SNAKE_CASE",
		testSnakeCase:          "snake_case",
		testTitleCase:          "Snake Case",
		testWordCase:           "snake case",
	},
	"kebab-case": {
		testCamelCase:          "kebabCase",
		testPascalCase:         "KebabCase",
		testDotCase:            "kebab.case",
		testKebabCase:          "kebab-case",
		testScreamingSnakeCase: "KEBAB_CASE",
		testSnakeCase:          "kebab_case",
		testTitleCase:          "Kebab Case",
		testWordCase:           "kebab case",
	},
	"IDOne": {
		testCamelCase:          "idOne",
		testPascalCase:         "IdOne",
		testDotCase:            "id.one",
		testKebabCase:          "id-one",
		testScreamingSnakeCase: "ID_ONE",
		testSnakeCase:          "id_one",
		testTitleCase:          "ID One",
		testWordCase:           "ID One",
	},
	"99two": {
		testCamelCase:          "99two",
		testPascalCase:         "99two",
		testDotCase:            "99two",
		testKebabCase:          "99two",
		testScreamingSnakeCase: "99TWO",
		testSnakeCase:          "99two",
		testTitleCase:          "99two",
		testWordCase:           "99two",
	},
	"99Two": {
		testCamelCase:          "99Two",
		testPascalCase:         "99Two",
		testDotCase:            "99.two",
		testKebabCase:          "99-two",
		testScreamingSnakeCase: "99_TWO",
		testSnakeCase:          "99_two",
		testTitleCase:          "99 Two",
		testWordCase:           "99 Two",
	},
	"interface{}": {
		testCamelCase:          "interface",
		testPascalCase:         "Interface",
		testDotCase:            "interface",
		testKebabCase:          "interface",
		testScreamingSnakeCase: "INTERFACE",
		testSnakeCase:          "interface",
		testTitleCase:          "Interface",
		testWordCase:           "interface",
	},
	"something$": {
		testCamelCase:          "something",
		testPascalCase:         "Something",
		testDotCase:            "something",
		testKebabCase:          "something",
		testScreamingSnakeCase: "SOMETHING",
		testSnakeCase:          "something",
		testTitleCase:          "Something",
		testWordCase:           "something",
	},
	"something$$$": {
		testCamelCase:          "something",
		testPascalCase:         "Something",
		testDotCase:            "something",
		testKebabCase:          "something",
		testScreamingSnakeCase: "SOMETHING",
		testSnakeCase:          "something",
		testTitleCase:          "Something",
		testWordCase:           "something",
	},
	"$prefixed": {
		testCamelCase:          "prefixed",
		testPascalCase:         "Prefixed",
		testDotCase:            "prefixed",
		testKebabCase:          "prefixed",
		testScreamingSnakeCase: "PREFIXED",
		testSnakeCase:          "prefixed",
		testTitleCase:          "Prefixed",
		testWordCase:           "prefixed",
	},
	"$$$prefixed": {
		testCamelCase:          "prefixed",
		testPascalCase:         "Prefixed",
		testDotCase:            "prefixed",
		testKebabCase:          "prefixed",
		testScreamingSnakeCase: "PREFIXED",
		testSnakeCase:          "prefixed",
		testTitleCase:          "Prefixed",
		testWordCase:           "prefixed",
	},
}

// TestJoinCase provides unit test coverage for Join()
func TestStandAloneMethods(t *testing.T) {
	for testText, tc := range testCases {
		for _, c := range testFunctions {
			t.Run(c+":"+testText, func(t *testing.T) {
				want, haveTest := tc[c]
				if haveTest {
					got := fnMap[c](testText)
					assert.Equal(t, want, got, "given: '%s' want: '%s' got: '%s'", testText, want, got)
				} else {
					assert.Fail(t, "Test case missing", "No test case for method '%s' for input '%s'", c, testText)
				}
			})
		}
	}
}

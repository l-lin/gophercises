package link

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParse_godoc_example(t *testing.T) {
	s := `
<p>
	Links:
</p>
<ul>
	<li><a href="foo">Foo</a>
	<li><a href="/bar/baz">BarBaz</a>
</ul>`
	links, err := Parse(strings.NewReader(s))
	if err != nil {
		t.Fatalf("Error when parsing: %v", err)
	}
	if len(links) != 2 {
		t.Errorf("Could not parse all 2 links, got %d links instead", len(links))
	}
	expected := []*Link{
		&Link{"foo", []string{"Foo"}},
		&Link{"/bar/baz", []string{"BarBaz"}},
	}
	for i := 0; i < 2; i++ {
		if links[i].Href != expected[i].Href || links[i].GetText() != expected[i].GetText() {
			t.Errorf("Parsed link = %v, want = %v", links[0], expected)
		}
	}
}

func TestParse_exercise_files(t *testing.T) {
	var tests = []struct {
		name  string
		wants []expected
		given string
	}{
		{
			"ex1.html - simple exercise",
			[]expected{
				expected{"/other-page", "A link to another page"},
			},
			"ex1.html",
		},
		{
			"ex2.html - nested nodes",
			[]expected{
				expected{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
				expected{"https://github.com/gophercises", "Gophercises is on Github !"},
			},
			"ex2.html",
		},
		{
			"ex0.html - nested nodes",
			[]expected{
				expected{"/dog", "Something in a span Text not in a span Bold text!"},
			},
			"ex0.html",
		},
		{
			"ex3.html - real life webpage",
			[]expected{
				expected{"#", "Login"},
				expected{"/lost", "Lost? Need help?"},
				expected{"https://twitter.com/marcusolsson", "@marcusolsson"},
			},
			"ex3.html",
		},
		{
			"ex4.html - comment inside",
			[]expected{
				expected{"/dog-cat", "dog cat"},
			},
			"ex4.html",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			content, err := ioutil.ReadFile(tt.given)
			if err != nil {
				t.Fatalf("Could not read file %s. Error was: %v", tt.given, err)
			}
			actual, err := Parse(bytes.NewReader(content))
			if err != nil {
				t.Fatalf("Could not parse file %s. Error was: %v", tt.given, err)
			}
			if len(actual) != len(tt.wants) {
				t.Errorf("(%s): expected %d links, actual %d links", tt.given, len(tt.wants), len(actual))
			}
			for i := 0; i < len(tt.wants); i++ {
				if actual[i].Href != tt.wants[i].href || actual[i].GetText() != tt.wants[i].text {
					t.Errorf("(%s):\nExpected \n%s\nActual \n%s\n", tt.given, tt.wants[i], actual[i])
				}

			}
		})
	}
}

type expected struct {
	href, text string
}

func (e expected) String() string {
	return fmt.Sprintf(`Link{
  Href: "%s",
  Text: "%s",
}`, e.href, e.text)
}

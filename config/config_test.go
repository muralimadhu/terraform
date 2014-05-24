package config

import (
	"path/filepath"
	"strings"
	"testing"
)

// This is the directory where our test fixtures are.
const fixtureDir = "./test-fixtures"

func TestConfigResourceGraph(t *testing.T) {
	c, err := Load(filepath.Join(fixtureDir, "resource_graph.tf"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	graph := c.ResourceGraph()
	if err := graph.Validate(); err != nil {
		t.Fatalf("err: %s", err)
	}

	actual := strings.TrimSpace(graph.String())
	expected := resourceGraphValue

	if actual != strings.TrimSpace(expected) {
		t.Fatalf("bad:\n%s", actual)
	}
}

func TestNewResourceVariable(t *testing.T) {
	v, err := NewResourceVariable("foo.bar.baz")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if v.Type != "foo" {
		t.Fatalf("bad: %#v", v)
	}
	if v.Name != "bar" {
		t.Fatalf("bad: %#v", v)
	}
	if v.Field != "baz" {
		t.Fatalf("bad: %#v", v)
	}

	if v.FullKey() != "foo.bar.baz" {
		t.Fatalf("bad: %#v", v)
	}
}

func TestNewUserVariable(t *testing.T) {
	v, err := NewUserVariable("var.bar")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if v.Name != "bar" {
		t.Fatalf("bad: %#v", v.Name)
	}
	if v.FullKey() != "var.bar" {
		t.Fatalf("bad: %#v", v)
	}
}

const resourceGraphValue = `
root: root
  root -> aws_security_group.firewall
  root -> aws_instance.web
aws_security_group.firewall
aws_instance.web
  aws_instance.web -> aws_security_group.firewall
root
  root -> aws_security_group.firewall
  root -> aws_instance.web
`
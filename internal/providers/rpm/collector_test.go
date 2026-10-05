package rpm

import (
	"errors"
	"reflect"
	"testing"

	"fedora-state/internal/model"
)

func TestParseRPMObjects(t *testing.T) {
	input := "" +
		"bash\t0\t5.3.0\t2.fc44\tx86_64\t100\tFedora Project\tFedora Project\tbash-5.3.0-2.fc44.src.rpm\n" +
		"example\t2\t1.4.0\t3.fc44\tnoarch\t200\tExample Vendor\tExample Packager\texample-1.4.0-3.fc44.src.rpm\n"

	got := parseRPMObjects(input)

	want := []model.Object{
		{
			Name:     "bash",
			Type:     "package",
			Identity: "rpm:bash:0:5.3.0-2.fc44:x86_64",
			Version:  "5.3.0-2.fc44",
			Evidence: []model.Evidence{
				{Type: "package_database", Value: "rpm"},
				{Type: "architecture", Value: "x86_64"},
				{Type: "install_time", Value: "100"},
				{Type: "vendor", Value: "Fedora Project"},
				{Type: "packager", Value: "Fedora Project"},
				{Type: "source_package", Value: "bash-5.3.0-2.fc44.src.rpm"},
			},
		},
		{
			Name:     "example",
			Type:     "package",
			Identity: "rpm:example:2:1.4.0-3.fc44:noarch",
			Version:  "2:1.4.0-3.fc44",
			Evidence: []model.Evidence{
				{Type: "package_database", Value: "rpm"},
				{Type: "architecture", Value: "noarch"},
				{Type: "install_time", Value: "200"},
				{Type: "vendor", Value: "Example Vendor"},
				{Type: "packager", Value: "Example Packager"},
				{Type: "source_package", Value: "example-1.4.0-3.fc44.src.rpm"},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("objects differ\n got: %#v\nwant: %#v", got, want)
	}
}

func TestParseRPMObjectsSkipsMalformedRecords(t *testing.T) {
	input := "" +
		"good\t0\t1.0\t1\tx86_64\t100\tVendor\tPackager\tgood-1.0-1.src.rpm\n" +
		"broken\n" +
		"\t0\t1.0\t1\tx86_64\t100\tVendor\tPackager\tbad.src.rpm\n"

	got := parseRPMObjects(input)

	if len(got) != 1 || got[0].Name != "good" {
		t.Fatalf("unexpected objects: %#v", got)
	}
}

func TestRPMCollectorUsesPackageDatabaseQuery(t *testing.T) {
	var gotName string
	var gotArgs []string

	collector := RPMCollector{
		command: func(name string, args ...string) ([]byte, error) {
			gotName = name
			gotArgs = append([]string(nil), args...)

			return []byte(
				"bash\t0\t5.3.0\t2.fc44\tx86_64\t100\tFedora Project\tFedora Project\tbash-5.3.0-2.fc44.src.rpm\n",
			), nil
		},
	}

	objects := collector.Collect().Objects

	if gotName != "rpm" {
		t.Fatalf("command=%q want rpm", gotName)
	}

	wantArgs := []string{
		"-qa",
		"--qf",
		"%{NAME}\\t%{EPOCHNUM}\\t%{VERSION}\\t%{RELEASE}\\t%{ARCH}\\t%{INSTALLTIME}\\t%{VENDOR}\\t%{PACKAGER}\\t%{SOURCERPM}\\n",
	}

	if !reflect.DeepEqual(gotArgs, wantArgs) {
		t.Fatalf("args differ\n got: %#v\nwant: %#v", gotArgs, wantArgs)
	}

	if len(objects) != 1 || objects[0].Identity != "rpm:bash:0:5.3.0-2.fc44:x86_64" {
		t.Fatalf("unexpected objects: %#v", objects)
	}
}

func TestRPMCollectorReturnsNoObjectsWhenRPMUnavailable(t *testing.T) {
	collector := RPMCollector{
		command: func(string, ...string) ([]byte, error) {
			return nil, errors.New("rpm unavailable")
		},
	}

	if got := collector.Collect(); got.Err == nil {
		t.Fatalf("expected collection error, got: %#v", got)
	}
}

package gopom

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

var p Project

func init() {
	err := xml.Unmarshal([]byte(examplePom), &p)
	if err != nil {
		fmt.Println("unable to parse xml ", err)
		os.Exit(1)
		return
	}
}

func TestParsing_Parent(t *testing.T) {
	if p.Parent.ArtifactID != "test-application" {
		t.Error("Parent.ArtifactID: expected 'test-application', got: " + p.Parent.ArtifactID)
	}
	if p.Parent.GroupID != "com.test" {
		t.Error("Parent.GroupID: expected 'com.test', got: " + p.Parent.GroupID)
	}
	if p.Parent.Version != "1.0.0" {
		t.Error("Parent.Version: expected '1.0.0', got: " + p.Parent.Version)
	}
	if p.Parent.RelativePath != "../pom.xml" {
		t.Error("Parent.RelativePath: expected '../pom.xml', got: " + p.Parent.RelativePath)
	}
}

func TestParsing_Project(t *testing.T) {
	if p.GroupID != "com.test" {
		t.Error("GroupID: expected 'com.test', got: " + p.GroupID)
	}
	if p.ArtifactID != "test-application" {
		t.Error("ArtifactID: expected 'test-application', got: " + p.ArtifactID)
	}
	if p.Version != "1.0.0" {
		t.Error("Version: expected '1.0.0', got: " + p.Version)
	}
	if p.Packaging != "war" {
		t.Error("Packaging: expected 'war', got: " + p.Packaging)
	}
	if p.Name != "gopom-test" {
		t.Error("Name: expected 'gopom-test', got: " + p.Name)
	}
	if p.Description != "pom parser in golang" {
		t.Error("Description: expected 'pom parser in golang', got: " + p.Description)
	}
	if p.URL != "testUrl" {
		t.Error("URL: expected 'testUrl', got: " + p.URL)
	}
	if p.InceptionYear != "2020" {
		t.Error("InceptionYear: expected '2020', got: " + p.InceptionYear)
	}
	if p.ModelVersion != "4.0.0" {
		t.Error("ModelVersion: expected '4.0.0', got: " + p.ModelVersion)
	}

	if len(p.Modules) != 2 {
		t.Error("Modules: expected len==2, got: " + strconv.Itoa(len(p.Modules)))
	} else {
		if p.Modules[0] != "module1" {
			t.Error("Modules[0]: expected 'module1', got: " + p.Modules[0])
		}
		if p.Modules[1] != "module2" {
			t.Error("Modules[1]: expected 'module1', got: " + p.Modules[1])
		}
	}
}

func TestParsing_Organization(t *testing.T) {
	if p.Organization.Name != "name" {
		t.Error("Organization.Name: expected 'name', got: " + p.Organization.Name)
	}
	if p.Organization.URL != "url" {
		t.Error("Organization.URL: expected 'url', got: " + p.Organization.URL)
	}
}

func TestParsing_Licenses(t *testing.T) {
	if len(p.Licenses) != 1 {
		t.Error("Licenses: expected len==1, got: " + strconv.Itoa(len(p.Licenses)))
	}
	if p.Licenses[0].Name != "name" {
		t.Error("Licenses.Name: expected 'name', got: " + p.Licenses[0].Name)
	}
	if p.Licenses[0].URL != "url" {
		t.Error("Licenses.URL: expected 'url', got: " + p.Licenses[0].URL)
	}
	if p.Licenses[0].Distribution != "manual" {
		t.Error("Licenses.Distribution: expected 'manual', got: " + p.Licenses[0].Distribution)
	}
	if p.Licenses[0].Comments != "comments" {
		t.Error("Licenses.Comments: expected 'comments', got: " + p.Licenses[0].Comments)
	}
}

func TestParsing_Developers(t *testing.T) {
	if len(p.Developers) != 1 {
		t.Error("Developers: expected len==1, got: " + strconv.Itoa(len(p.Developers)))
	}

	d := p.Developers[0]
	if d.ID != "id" {
		t.Error("Developers[0].ID: expected 'id', got: " + d.ID)
	}
	if d.Name != "name" {
		t.Error("Developers[0].Name: expected 'name', got: " + d.Name)
	}
	if d.Email != "email" {
		t.Error("Developers[0].Email: expected 'email', got: " + d.Email)
	}
	if d.URL != "url" {
		t.Error("Developers[0].URL: expected 'url', got: " + d.URL)
	}
	if d.Organization != "organization" {
		t.Error("Developers[0].Organization: expected 'organization', got: " + d.Organization)
	}
	if d.OrganizationURL != "organizationUrl" {
		t.Error("Developers[0].OrganizationUrl: expected 'organizationUrl', got: " + d.OrganizationURL)
	}
	if d.Timezone != "+1" {
		t.Error("Developers[0].Timezone: expected '+1', got: " + d.Timezone)
	}
	if len(d.Roles) != 2 {
		t.Error("Developers: expected len==2, got: " + strconv.Itoa(len(d.Roles)))
	}
	if d.Roles[0] != "role1" {
		t.Error("Developers[0].Roles[0]: expected 'role1', got: " + d.Roles[0])
	}
	if d.Roles[1] != "role2" {
		t.Error("Developers[0].Roles[1]: expected 'role2', got: " + d.Roles[1])
	}
}

func TestParse_Contributors(t *testing.T) {
	if len(p.Contributors) != 1 {
		t.Error("Contributors: expected len==1, got: " + strconv.Itoa(len(p.Contributors)))
	}

	c := p.Contributors[0]
	if c.Name != "name" {
		t.Error("Contributors[0].Name: expected 'name', got: " + c.Name)
	}
	if c.Email != "email" {
		t.Error("Contributors[0].Email: expected 'email', got: " + c.Email)
	}
	if c.URL != "url" {
		t.Error("Contributors[0].URL: expected 'url', got: " + c.URL)
	}
	if c.Organization != "organization" {
		t.Error("Contributors[0].Organization: expected 'organization', got: " + c.Organization)
	}
	if c.OrganizationURL != "organizationUrl" {
		t.Error("Contributors[0].OrganizationUrl: expected 'organizationUrl', got: " + c.OrganizationURL)
	}
	if c.Timezone != "+1" {
		t.Error("Contributors[0].Timezone: expected '+1', got: " + c.Timezone)
	}
	if len(c.Roles) != 2 {
		t.Error("Contributors: expected len==2, got: " + strconv.Itoa(len(c.Roles)))
	}
	if c.Roles[0] != "role1" {
		t.Error("Contributors[0].Roles[0]: expected 'role1', got: " + c.Roles[0])
	}
	if c.Roles[1] != "role2" {
		t.Error("Developers[0].Roles[1]: expected 'role2', got: " + c.Roles[1])
	}
}

func TestParse_MailingLists(t *testing.T) {
	if len(p.MailingLists) != 1 {
		t.Error("MailingLists: expected len==1, got: " + strconv.Itoa(len(p.MailingLists)))
	}

	m := p.MailingLists[0]
	if m.Name != "name" {
		t.Error("MailingLists[0].Name: expected 'name', got: " + m.Name)
	}
	if m.Subscribe != "subscribe" {
		t.Error("MailingLists[0].Subscribe: expected 'subscribe', got: " + m.Subscribe)
	}
	if m.Unsubscribe != "unsubscribe" {
		t.Error("MailingLists[0].Unsubscribe: expected 'unsubscribe', got: " + m.Unsubscribe)
	}
	if m.Post != "post" {
		t.Error("MailingLists[0].Post: expected 'post', got: " + m.Post)
	}
	if m.Archive != "archive" {
		t.Error("MailingLists[0].Archive: expected 'archive', got: " + m.Archive)
	}

	if len(m.OtherArchives) != 2 {
		t.Error("MailingLists.OtherArchives: expected len==2, got: " + strconv.Itoa(len(m.OtherArchives)))
	}
	if m.OtherArchives[0] != "archive1" {
		t.Error("MailingLists[0].OtherArchives[0]: expected 'archive1', got: " + m.OtherArchives[0])
	}
	if m.OtherArchives[1] != "archive2" {
		t.Error("MailingLists[0].OtherArchives[1]: expected 'archive2', got: " + m.OtherArchives[1])
	}
}

func TestParsing_Prerequisites(t *testing.T) {
	if p.Prerequisites.Maven != "2.0" {
		t.Error("Prerequisites.Maven: expected '2.0', got: " + p.Prerequisites.Maven)
	}
}

func TestParsing_SCM(t *testing.T) {
	if p.SCM.URL != "url" {
		t.Error("SCM.URL: expected 'url', got: " + p.SCM.URL)
	}
	if p.SCM.Connection != "connection" {
		t.Error("SCM.Connection: expected 'connection', got: " + p.SCM.Connection)
	}
	if p.SCM.DeveloperConnection != "developerConnection" {
		t.Error("SCM.DeveloperConnection: expected 'developerConnection', got: " + p.SCM.DeveloperConnection)
	}
	if p.SCM.Tag != "tag" {
		t.Error("SCM.Tag: expected 'tag', got: " + p.SCM.Tag)
	}
}

func TestParsing_IssueManagement(t *testing.T) {
	if p.IssueManagement.URL != "url" {
		t.Error("IssueManagement.URL: expected 'url', got: " + p.IssueManagement.URL)
	}
	if p.IssueManagement.System != "system" {
		t.Error("IssueManagement.System: expected 'system', got: " + p.IssueManagement.System)
	}
}

func TestParsing_CIManagement(t *testing.T) {
	if p.CIManagement.System != "system" {
		t.Error("CIManagement.System: expected 'system', got: " + p.CIManagement.System)
	}
	if p.CIManagement.URL != "url" {
		t.Error("CIManagement.URL: expected 'url', got: " + p.CIManagement.URL)
	}
}

func TestParsing_Notifiers(t *testing.T) {
	if len(p.CIManagement.Notifiers) != 1 {
		t.Error("CIManagement.Notifiers: expected len==1, got: " + strconv.Itoa(len(p.CIManagement.Notifiers)))
	}
	n := p.CIManagement.Notifiers[0]
	if n.Type != "type" {
		t.Error("CIManagement.Notifiers.Type: expected 'type', got: " + n.Type)
	}
	if n.Address != "address" {
		t.Error("CIManagement.Notifiers.Address: expected 'type', got: " + n.Address)
	}
	if n.SendOnError != true {
		t.Error("CIManagement.Notifiers.SendOnError: expected 'true', got: " + strconv.FormatBool(n.SendOnError))
	}
	if n.SendOnFailure != true {
		t.Error("CIManagement.Notifiers.SendOnFailure: expected 'true', got: " + strconv.FormatBool(n.SendOnFailure))
	}
	if n.SendOnSuccess != true {
		t.Error("CIManagement.Notifiers.SendOnSuccess: expected 'true', got: " + strconv.FormatBool(n.SendOnSuccess))
	}
	if n.SendOnWarning != true {
		t.Error("CIManagement.Notifiers.SendOnWarning: expected 'true', got: " + strconv.FormatBool(n.SendOnWarning))
	}
}

func TestParsing_DistributionManagement(t *testing.T) {
	assert.Equal(t, "downloadUrl", p.DistributionManagement.DownloadURL)
	assert.Equal(t, "status", p.DistributionManagement.Status)
	assert.Equal(t, "name", p.DistributionManagement.Repository.Name)
	assert.Equal(t, "url", p.DistributionManagement.Repository.URL)
	assert.Equal(t, "layout", p.DistributionManagement.Repository.Layout)
	assert.Equal(t, true, p.DistributionManagement.Repository.UniqueVersion)
	assert.Equal(t, "id", p.DistributionManagement.Repository.ID)
	r := p.DistributionManagement.Repository.Releases
	assert.Equal(t, "checksumPolicy", r.ChecksumPolicy)
	assert.Equal(t, "enabled", r.Enabled)
	assert.Equal(t, "updatePolicy", r.UpdatePolicy)
	s := p.DistributionManagement.Repository.Snapshots
	assert.Equal(t, "checksumPolicy", s.ChecksumPolicy)
	assert.Equal(t, "enabled", s.Enabled)
	assert.Equal(t, "updatePolicy", s.UpdatePolicy)

	sr := p.DistributionManagement.SnapshotRepository
	assert.Equal(t, "name", sr.Name)
	assert.Equal(t, "url", sr.URL)
	assert.Equal(t, "layout", sr.Layout)
	assert.Equal(t, true, sr.UniqueVersion)
	assert.Equal(t, "id", sr.ID)
	r = sr.Releases
	assert.Equal(t, "checksumPolicy", r.ChecksumPolicy)
	assert.Equal(t, "enabled", r.Enabled)
	assert.Equal(t, "updatePolicy", r.UpdatePolicy)
	s = sr.Snapshots
	assert.Equal(t, "checksumPolicy", s.ChecksumPolicy)
	assert.Equal(t, "enabled", s.Enabled)
	assert.Equal(t, "updatePolicy", s.UpdatePolicy)

	rel := p.DistributionManagement.Relocation
	assert.Equal(t, "version", rel.Version)
	assert.Equal(t, "artifactId", rel.ArtifactID)
	assert.Equal(t, "groupId", rel.GroupID)
	assert.Equal(t, "message", rel.Message)

	site := p.DistributionManagement.Site
	assert.Equal(t, "id", site.ID)
	assert.Equal(t, "url", site.URL)
	assert.Equal(t, "name", site.Name)
}

func TestParsing_DependencyManagement(t *testing.T) {
	assert.Equal(t, 1, len(p.DependencyManagement.Dependencies))
	d := p.DependencyManagement.Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)
}

func TestParsing_Dependencies(t *testing.T) {
	assert.Equal(t, 2, len(p.Dependencies))
	d := p.Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)
}

func TestParsing_Repositories(t *testing.T) {
	assert.Equal(t, 1, len(p.Repositories))
	r := p.Repositories[0]
	assert.Equal(t, "id", r.ID)
	assert.Equal(t, "name", r.Name)
	assert.Equal(t, "url", r.URL)
	assert.Equal(t, "layout", r.Layout)
	assert.Equal(t, "enabled", r.Releases.Enabled)
	assert.Equal(t, "checksumPolicy", r.Releases.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", r.Releases.UpdatePolicy)
	assert.Equal(t, "enabled", r.Snapshots.Enabled)
	assert.Equal(t, "checksumPolicy", r.Snapshots.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", r.Snapshots.UpdatePolicy)
}

func TestParsing_PluginRepositories(t *testing.T) {
	assert.Equal(t, 1, len(p.PluginRepositories))
	pr := p.PluginRepositories[0]
	assert.Equal(t, "id", pr.ID)
	assert.Equal(t, "name", pr.Name)
	assert.Equal(t, "url", pr.URL)
	assert.Equal(t, "layout", pr.Layout)
	assert.Equal(t, "enabled", pr.Releases.Enabled)
	assert.Equal(t, "checksumPolicy", pr.Releases.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", pr.Releases.UpdatePolicy)
	assert.Equal(t, "enabled", pr.Snapshots.Enabled)
	assert.Equal(t, "checksumPolicy", pr.Snapshots.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", pr.Snapshots.UpdatePolicy)
}

func TestParsing_Build(t *testing.T) {
	b := p.Build
	assert.Equal(t, "sourceDirectory", b.SourceDirectory)
	assert.Equal(t, "scriptSourceDirectory", b.ScriptSourceDirectory)
	assert.Equal(t, "testSourceDirectory", b.TestSourceDirectory)
	assert.Equal(t, "outputDirectory", b.OutputDirectory)
	assert.Equal(t, "testOutputDirectory", b.TestOutputDirectory)
	assert.Equal(t, "directory", b.Directory)
	assert.Equal(t, "defaultGoal", b.DefaultGoal)
	assert.Equal(t, "finalName", b.FinalName)
	assert.Equal(t, 1, len(b.Filters))
	assert.Equal(t, "filter1", b.Filters[0])

	assert.Equal(t, 1, len(b.Extensions))
	assert.Equal(t, "groupId", b.Extensions[0].GroupID)
	assert.Equal(t, "artifactId", b.Extensions[0].ArtifactID)
	assert.Equal(t, "version", b.Extensions[0].Version)

	assert.Equal(t, 1, len(b.Resources))
	assert.Equal(t, "targetPath", b.Resources[0].TargetPath)
	assert.Equal(t, "filtering", b.Resources[0].Filtering)
	assert.Equal(t, "directory", b.Resources[0].Directory)
	assert.Equal(t, 1, len(b.Resources[0].Includes))
	assert.Equal(t, "include", b.Resources[0].Includes[0])
	assert.Equal(t, 1, len(b.Resources[0].Excludes))
	assert.Equal(t, "exclude", b.Resources[0].Excludes[0])

	assert.Equal(t, 1, len(b.TestResources))
	assert.Equal(t, "targetPath", b.TestResources[0].TargetPath)
	assert.Equal(t, "filtering", b.TestResources[0].Filtering)
	assert.Equal(t, "directory", b.TestResources[0].Directory)
	assert.Equal(t, 1, len(b.TestResources[0].Includes))
	assert.Equal(t, "include", b.TestResources[0].Includes[0])
	assert.Equal(t, 1, len(b.TestResources[0].Excludes))
	assert.Equal(t, "exclude", b.TestResources[0].Excludes[0])

	pl := b.PluginManagement.Plugins
	assert.Equal(t, 1, len(pl))
	assert.Equal(t, "groupId", pl[0].GroupID)
	assert.Equal(t, "artifactId", pl[0].ArtifactID)
	assert.Equal(t, "version", pl[0].Version)
	assert.Equal(t, "extensions", pl[0].Extensions)
	assert.Equal(t, 1, len(pl[0].Executions))
	assert.Equal(t, "id", pl[0].Executions[0].ID)
	assert.Equal(t, "phase", pl[0].Executions[0].Phase)
	assert.Equal(t, 1, len(pl[0].Executions[0].Goals))
	assert.Equal(t, "goal", pl[0].Executions[0].Goals[0])
	assert.Equal(t, "inherited", pl[0].Executions[0].Inherited)

	assert.Equal(t, 1, len(pl[0].Dependencies))
	d := pl[0].Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)

	pl = b.Plugins
	assert.Equal(t, 1, len(pl))
	assert.Equal(t, "groupId", pl[0].GroupID)
	assert.Equal(t, "artifactId", pl[0].ArtifactID)
	assert.Equal(t, "version", pl[0].Version)
	assert.Equal(t, "extensions", pl[0].Extensions)
	assert.Equal(t, 1, len(pl[0].Executions))
	assert.Equal(t, "id", pl[0].Executions[0].ID)
	assert.Equal(t, "phase", pl[0].Executions[0].Phase)
	assert.Equal(t, 1, len(pl[0].Executions[0].Goals))
	assert.Equal(t, "goal", pl[0].Executions[0].Goals[0])
	assert.Equal(t, "inherited", pl[0].Executions[0].Inherited)

	assert.Equal(t, 1, len(pl[0].Dependencies))
	d = pl[0].Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)
}

func TestParsing_Reporting(t *testing.T) {
	assert.Equal(t, "excludeDefaults", p.Reporting.ExcludeDefaults)
	assert.Equal(t, "outputDirectory", p.Reporting.OutputDirectory)
	assert.Equal(t, "outputDirectory", p.Reporting.OutputDirectory)

	pl := p.Reporting.Plugins
	assert.Equal(t, 1, len(pl))
	assert.Equal(t, "groupId", pl[0].GroupID)
	assert.Equal(t, "artifactId", pl[0].ArtifactID)
	assert.Equal(t, "version", pl[0].Version)
	assert.Equal(t, 1, len(pl[0].ReportSets))
	assert.Equal(t, "id", pl[0].ReportSets[0].ID)
	assert.Equal(t, 1, len(pl[0].ReportSets[0].Reports))
	assert.Equal(t, "report", pl[0].ReportSets[0].Reports[0])
	assert.Equal(t, "inherited", pl[0].ReportSets[0].Inherited)
}

func TestParsing_Profiles(t *testing.T) {
	assert.Equal(t, 1, len(p.Profiles))
	assert.Equal(t, "id", p.Profiles[0].ID)
	assert.Equal(t, true, p.Profiles[0].Activation.ActiveByDefault)
	assert.Equal(t, "jdk", p.Profiles[0].Activation.JDK)
	assert.Equal(t, "name", p.Profiles[0].Activation.OS.Name)
	assert.Equal(t, "family", p.Profiles[0].Activation.OS.Family)
	assert.Equal(t, "arch", p.Profiles[0].Activation.OS.Arch)
	assert.Equal(t, "version", p.Profiles[0].Activation.OS.Version)
	assert.Equal(t, "name", p.Profiles[0].Activation.Property.Name)
	assert.Equal(t, "value", p.Profiles[0].Activation.Property.Value)
	assert.Equal(t, "missing", p.Profiles[0].Activation.File.Missing)
	assert.Equal(t, "exists", p.Profiles[0].Activation.File.Exists)

	b := p.Profiles[0].Build
	assert.Equal(t, "directory", b.Directory)
	assert.Equal(t, "defaultGoal", b.DefaultGoal)
	assert.Equal(t, "finalName", b.FinalName)
	assert.Equal(t, 1, len(b.Filters))
	assert.Equal(t, "filter1", b.Filters[0])

	assert.Equal(t, 1, len(b.Resources))
	assert.Equal(t, "targetPath", b.Resources[0].TargetPath)
	assert.Equal(t, "filtering", b.Resources[0].Filtering)
	assert.Equal(t, "directory", b.Resources[0].Directory)
	assert.Equal(t, 1, len(b.Resources[0].Includes))
	assert.Equal(t, "include", b.Resources[0].Includes[0])
	assert.Equal(t, 1, len(b.Resources[0].Excludes))
	assert.Equal(t, "exclude", b.Resources[0].Excludes[0])

	assert.Equal(t, 1, len(b.TestResources))
	assert.Equal(t, "targetPath", b.TestResources[0].TargetPath)
	assert.Equal(t, "filtering", b.TestResources[0].Filtering)
	assert.Equal(t, "directory", b.TestResources[0].Directory)
	assert.Equal(t, 1, len(b.TestResources[0].Includes))
	assert.Equal(t, "include", b.TestResources[0].Includes[0])
	assert.Equal(t, 1, len(b.TestResources[0].Excludes))
	assert.Equal(t, "exclude", b.TestResources[0].Excludes[0])

	pl := b.PluginManagement.Plugins
	assert.Equal(t, 1, len(pl))
	assert.Equal(t, "groupId", pl[0].GroupID)
	assert.Equal(t, "artifactId", pl[0].ArtifactID)
	assert.Equal(t, "version", pl[0].Version)
	assert.Equal(t, "extensions", pl[0].Extensions)
	assert.Equal(t, 1, len(pl[0].Executions))
	assert.Equal(t, "id", pl[0].Executions[0].ID)
	assert.Equal(t, "phase", pl[0].Executions[0].Phase)
	assert.Equal(t, 1, len(pl[0].Executions[0].Goals))
	assert.Equal(t, "goal", pl[0].Executions[0].Goals[0])
	assert.Equal(t, "inherited", pl[0].Executions[0].Inherited)

	assert.Equal(t, 1, len(pl[0].Dependencies))
	d := pl[0].Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)

	pl = b.Plugins
	assert.Equal(t, 1, len(pl))
	assert.Equal(t, "groupId", pl[0].GroupID)
	assert.Equal(t, "artifactId", pl[0].ArtifactID)
	assert.Equal(t, "version", pl[0].Version)
	assert.Equal(t, "extensions", pl[0].Extensions)
	assert.Equal(t, 1, len(pl[0].Executions))
	assert.Equal(t, "id", pl[0].Executions[0].ID)
	assert.Equal(t, "phase", pl[0].Executions[0].Phase)
	assert.Equal(t, 1, len(pl[0].Executions[0].Goals))
	assert.Equal(t, "goal", pl[0].Executions[0].Goals[0])
	assert.Equal(t, "inherited", pl[0].Executions[0].Inherited)

	assert.Equal(t, 1, len(pl[0].Dependencies))
	d = pl[0].Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)

	assert.Equal(t, "module1", p.Profiles[0].Modules[0])

	dm := p.Profiles[0].DistributionManagement
	assert.Equal(t, "downloadUrl", dm.DownloadURL)
	assert.Equal(t, "status", dm.Status)
	assert.Equal(t, "name", dm.Repository.Name)
	assert.Equal(t, "url", dm.Repository.URL)
	assert.Equal(t, "layout", dm.Repository.Layout)
	assert.Equal(t, true, dm.Repository.UniqueVersion)
	assert.Equal(t, "id", dm.Repository.ID)
	r := dm.Repository.Releases
	assert.Equal(t, "checksumPolicy", r.ChecksumPolicy)
	assert.Equal(t, "enabled", r.Enabled)
	assert.Equal(t, "updatePolicy", r.UpdatePolicy)
	s := dm.Repository.Snapshots
	assert.Equal(t, "checksumPolicy", s.ChecksumPolicy)
	assert.Equal(t, "enabled", s.Enabled)
	assert.Equal(t, "updatePolicy", s.UpdatePolicy)

	sr := dm.SnapshotRepository
	assert.Equal(t, "name", sr.Name)
	assert.Equal(t, "url", sr.URL)
	assert.Equal(t, "layout", sr.Layout)
	assert.Equal(t, true, sr.UniqueVersion)
	assert.Equal(t, "id", sr.ID)
	r = sr.Releases
	assert.Equal(t, "checksumPolicy", r.ChecksumPolicy)
	assert.Equal(t, "enabled", r.Enabled)
	assert.Equal(t, "updatePolicy", r.UpdatePolicy)
	s = sr.Snapshots
	assert.Equal(t, "checksumPolicy", s.ChecksumPolicy)
	assert.Equal(t, "enabled", s.Enabled)
	assert.Equal(t, "updatePolicy", s.UpdatePolicy)

	rel := dm.Relocation
	assert.Equal(t, "version", rel.Version)
	assert.Equal(t, "artifactId", rel.ArtifactID)
	assert.Equal(t, "groupId", rel.GroupID)
	assert.Equal(t, "message", rel.Message)

	site := dm.Site
	assert.Equal(t, "id", site.ID)
	assert.Equal(t, "url", site.URL)
	assert.Equal(t, "name", site.Name)

	dMan := p.Profiles[0].DependencyManagement
	assert.Equal(t, 1, len(dMan.Dependencies))
	d = dMan.Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)

	d = p.Profiles[0].Dependencies[0]
	assert.Equal(t, "groupId", d.GroupID)
	assert.Equal(t, "artifactId", d.ArtifactID)
	assert.Equal(t, "version", d.Version)
	assert.Equal(t, "type", d.Type)
	assert.Equal(t, "classifier", d.Classifier)
	assert.Equal(t, 1, len(d.Exclusions))
	assert.Equal(t, "artifactId", d.Exclusions[0].ArtifactID)
	assert.Equal(t, "groupId", d.Exclusions[0].GroupID)
	assert.Equal(t, "optional", d.Optional)
	assert.Equal(t, "scope", d.Scope)
	assert.Equal(t, "systemPath", d.SystemPath)

	rep := p.Profiles[0].Repositories[0]
	assert.Equal(t, "id", rep.ID)
	assert.Equal(t, "name", rep.Name)
	assert.Equal(t, "url", rep.URL)
	assert.Equal(t, "layout", rep.Layout)
	assert.Equal(t, "enabled", rep.Releases.Enabled)
	assert.Equal(t, "checksumPolicy", rep.Releases.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", rep.Releases.UpdatePolicy)
	assert.Equal(t, "enabled", rep.Snapshots.Enabled)
	assert.Equal(t, "checksumPolicy", rep.Snapshots.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", rep.Snapshots.UpdatePolicy)

	pluRep := p.Profiles[0].PluginRepositories[0]
	assert.Equal(t, "id", pluRep.ID)
	assert.Equal(t, "name", pluRep.Name)
	assert.Equal(t, "url", pluRep.URL)
	assert.Equal(t, "layout", pluRep.Layout)
	assert.Equal(t, "enabled", pluRep.Releases.Enabled)
	assert.Equal(t, "checksumPolicy", pluRep.Releases.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", pluRep.Releases.UpdatePolicy)
	assert.Equal(t, "enabled", pluRep.Snapshots.Enabled)
	assert.Equal(t, "checksumPolicy", pluRep.Snapshots.ChecksumPolicy)
	assert.Equal(t, "updatePolicy", pluRep.Snapshots.UpdatePolicy)

	reporting := p.Profiles[0].Reporting
	assert.Equal(t, "excludeDefaults", reporting.ExcludeDefaults)
	assert.Equal(t, "outputDirectory", reporting.OutputDirectory)
	assert.Equal(t, "outputDirectory", reporting.OutputDirectory)

	repPl := reporting.Plugins
	assert.Equal(t, 1, len(repPl))
	assert.Equal(t, "groupId", repPl[0].GroupID)
	assert.Equal(t, "artifactId", repPl[0].ArtifactID)
	assert.Equal(t, "version", repPl[0].Version)
	assert.Equal(t, "inherited", repPl[0].Inherited)
	assert.Equal(t, 1, len(repPl[0].ReportSets))
	assert.Equal(t, "id", repPl[0].ReportSets[0].ID)
	assert.Equal(t, 1, len(repPl[0].ReportSets[0].Reports))
	assert.Equal(t, "report", repPl[0].ReportSets[0].Reports[0])
	assert.Equal(t, "inherited", repPl[0].ReportSets[0].Inherited)
}

func Test_ParsingParentProperties(t *testing.T) {
	assert.Equal(t, 4, len(p.Properties.Entries))
	assert.Equal(t, "value", p.Properties.Entries["key"])
	assert.Equal(t, "value2", p.Properties.Entries["key2"])
	assert.Equal(t, "value3", p.Properties.Entries["key3"])
}

func Test_ParsingDeveloperProperties(t *testing.T) {
	assert.Equal(t, 3, len(p.Developers[0].Properties.Entries))
	assert.Equal(t, "value", p.Properties.Entries["key"])
	assert.Equal(t, "value2", p.Properties.Entries["key2"])
	assert.Equal(t, "value3", p.Properties.Entries["key3"])
}

func Test_ParsingContributorProperties(t *testing.T) {
	assert.Equal(t, 3, len(p.Contributors[0].Properties.Entries))
	assert.Equal(t, "value", p.Properties.Entries["key"])
	assert.Equal(t, "value2", p.Properties.Entries["key2"])
	assert.Equal(t, "value3", p.Properties.Entries["key3"])
}

func Test_ParsingProfileProperties(t *testing.T) {
	assert.Equal(t, 3, len(p.Profiles[0].Properties.Entries))
	assert.Equal(t, "value", p.Properties.Entries["key"])
	assert.Equal(t, "value2", p.Properties.Entries["key2"])
	assert.Equal(t, "value3", p.Properties.Entries["key3"])
}

func Test_ParsingNotifierConfigurations(t *testing.T) {
	assert.Equal(t, 3, len(p.CIManagement.Notifiers[0].Configuration.Entries))
	assert.Equal(t, "value", p.Properties.Entries["key"])
	assert.Equal(t, "value2", p.Properties.Entries["key2"])
	assert.Equal(t, "value3", p.Properties.Entries["key3"])
}

func TestNonExpandedVariables(t *testing.T) {
	dep := p.Dependencies[1]
	assert.NotNil(t, dep)
	assert.Equal(t, "com.group.test", dep.GroupID)
	assert.Equal(t, "artifact-id", dep.ArtifactID)
	assert.Equal(t, "${groupId.artifactId.version}", dep.Version)
	assert.Equal(t, "com/group/test/artifact-id/${groupId.artifactId.version}/artifact-id-${groupId.artifactId.version}", dep.NormalizedName())
	assert.Equal(t, "com/group/test/artifact-id/${groupId.artifactId.version}/artifact-id-${groupId.artifactId.version}.jar", dep.ToURL())
}

func TestExpandingProjectProperty(t *testing.T) {
	dep := p.Dependencies[1]
	assert.NotNil(t, dep)
	assert.Equal(t, "com/group/test/artifact-id/1.0.0/artifact-id-1.0.0", p.ExpandVariable(dep.NormalizedName()))
	assert.Equal(t, "com/group/test/artifact-id/1.0.0/artifact-id-1.0.0.jar", p.ExpandVariable(dep.ToURL()))
}




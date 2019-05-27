package cmd

import "testing"

func (p *MultiplyCommand) TestExecute(t *testing.T) {

	total := p.Execute([]string{"0", "0", "2", "6"})

	if total != "6" {
		t.Errorf("Multi was incorrect, got: %s, want: %d.", total, 6)
	}

}

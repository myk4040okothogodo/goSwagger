package data

import "testing"

func TestChecksValidation(t *testing.T){
  p := &Product{
    Name: "nics",
    Price: 1.90,
    SKU: "ans-absjr-swhrh"
  }

  err := p.Validate()
  if err != nil {
    t.Fatal(err)
  }

}

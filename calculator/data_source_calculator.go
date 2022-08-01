package calculator

import (
	"context"
	"fmt"
	"github.com/Looty/go-calculator/calculator"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCalculator() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCalculatorRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
				ForceNew: true,
			},
			"a": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"b": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"function": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if !IsValidFunction(v) {
						errs = append(errs, fmt.Errorf("%q must be [add, sub, mul, div], got: '%s'", key, v))
					}
					return
				},
			},
			"result": {
				Type:     schema.TypeInt,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func dataSourceCalculatorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type

	var diags diag.Diagnostics

	a := d.Get("a").(int)
	b := d.Get("b").(int)
	function := d.Get("function").(string)

	result := 0
	var err error

	switch function {
	case "add":
		result, err = calculator.Add(a, b)
	case "sub":
		result, err = calculator.Sub(a, b)
	case "mul":
		result, err = calculator.Mul(a, b)
	case "div":
		result, err = calculator.Div(a, b)
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("result", result); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func IsValidFunction(function string) bool {
	switch function {
	case
		"add",
		"sub",
		"mul",
		"div":
		return true
	}
	return false
}

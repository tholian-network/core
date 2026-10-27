package structs

import "testing"

func TestNewProduct(t *testing.T) {

	product := NewProduct("software")

	if product.Name != "" {
		t.Errorf("NewProduct() Name = %q, expected empty", product.Name)
	}

	if product.Vendor != "any" {
		t.Errorf("NewProduct() Vendor = %q, expected %q", product.Vendor, "any")
	}

	if product.Product != "any" {
		t.Errorf("NewProduct() Product = %q, expected %q", product.Product, "any")
	}

	if product.Version != "any" {
		t.Errorf("NewProduct() Version = %q, expected %q", product.Version, "any")
	}

	if product.Type != "software" {
		t.Errorf("NewProduct() Type = %q, expected %q", product.Type, "software")
	}

	if product.State != "invalid" {
		t.Errorf("NewProduct() State = %q, expected %q", product.State, "invalid")
	}

	if product.IsValid() == true {
		t.Fatal("a new product should be invalid")
	}

	types := []string{"hardware", "software", "system"}

	for _, typ := range types {

		value := NewProduct(typ)

		if value.Type != typ {
			t.Errorf("NewProduct(%q) Type = %q", typ, value.Type)
		}

	}

	invalid := NewProduct("invalid")

	if invalid.Type != "" {
		t.Errorf("NewProduct() accepted an invalid type: %q", invalid.Type)
	}

}

func TestProductSetters(t *testing.T) {

	product := NewProduct("software")

	product.SetVendor("  acme  ")
	product.SetProduct("  widget  ")
	product.SetVersion("  1.0  ")
	product.SetState("published")

	if product.Vendor != "acme" {
		t.Errorf("SetVendor() = %q, expected %q", product.Vendor, "acme")
	}

	if product.Product != "widget" {
		t.Errorf("SetProduct() = %q, expected %q", product.Product, "widget")
	}

	if product.Version != "10" {
		t.Errorf("SetVersion() = %q, expected %q", product.Version, "10")
	}

	if product.Name != "acme:widget:10" {
		t.Errorf("Name = %q, expected %q", product.Name, "acme:widget:10")
	}

	if product.State != "published" {
		t.Errorf("SetState() = %q, expected %q", product.State, "published")
	}

	if product.IsValid() == false {
		t.Fatal("a populated and published product should be valid")
	}

	product.SetType("invalid")
	product.SetState("invalid-state")

	if product.Type != "software" {
		t.Errorf("SetType() accepted an invalid type: %q", product.Type)
	}

	if product.State != "published" {
		t.Errorf("SetState() accepted an invalid state: %q", product.State)
	}

	product.SetVendor("")
	product.SetProduct("")
	product.SetVersion("")

	if product.Vendor != "acme" || product.Product != "widget" || product.Version != "10" {
		t.Errorf("Setters accepted empty values: %+v", product)
	}

	product.SetState("edited")

	if product.State != "edited" {
		t.Errorf("SetState() = %q, expected %q", product.State, "edited")
	}

	product.SetState("invalid")

	if product.State != "invalid" {
		t.Errorf("SetState() = %q, expected %q", product.State, "invalid")
	}

}

func TestProductWildcardDash(t *testing.T) {

	new_product := func() Product {

		product := NewProduct("software")
		product.SetVendor("acme")
		product.SetProduct("widget")
		product.SetVersion("1.0")

		return product

	}

	vendor_wildcards := []string{"*", "-"}

	for _, value := range vendor_wildcards {

		product := new_product()
		product.SetVendor(value)

		if product.Vendor != "" || product.Name != "" {
			t.Errorf("SetVendor(%q) = %q, name %q; expected empty values", value, product.Vendor, product.Name)
		}

	}

	product_wildcards := []string{"*", "-"}

	for _, value := range product_wildcards {

		product := new_product()
		product.SetProduct(value)

		if product.Product != "" || product.Name != "" {
			t.Errorf("SetProduct(%q) = %q, name %q; expected empty values", value, product.Product, product.Name)
		}

	}

	version_wildcards := []string{"any", "*", "-"}

	for _, value := range version_wildcards {

		product := new_product()
		product.SetVersion(value)

		if product.Version != "any" || product.Name != "" {
			t.Errorf("SetVersion(%q) = %q, name %q; expected any and empty name", value, product.Version, product.Name)
		}

	}

}

func TestProductToProductCPE23(t *testing.T) {

	cases := []struct {
		name        string
		value       string
		wantType    string
		wantVendor  string
		wantProduct string
		wantVersion string
		wantName    string
		wantValid   bool
	}{
		{
			"full",
			"cpe:2.3:a:vendor:product:1.0:beta:*:*:*:*:*:*:*",
			"software", "vendor", "product", "10-beta", "vendor:product:10-beta", true,
		},
		{
			"version only",
			"cpe:2.3:a:vendor:product:2.0",
			"software", "vendor", "product", "20", "vendor:product:20", true,
		},
		{
			"hardware wildcard",
			"cpe:2.3:h:vendor:product:*:*:*:*:*:*:*:*:*",
			"hardware", "vendor", "product", "any", "vendor:product:any", true,
		},
		{
			"system minimal",
			"cpe:2.3:o:vendor:product",
			"system", "vendor", "product", "any", "vendor:product:any", true,
		},
		{
			"architecture suffix",
			"cpe:2.3:a:vendor:product:1.0_x86_64:*:*:*:*:*:*:*",
			"software", "vendor", "product", "10", "vendor:product:10", true,
		},
		{
			"architecture prefix",
			"cpe:2.3:a:vendor:product:x86_64_1.0:*:*:*:*:*:*:*",
			"software", "vendor", "product", "10", "vendor:product:10", true,
		},
		{
			"all wildcards",
			"cpe:2.3:a:*:*:*:*:*:*:*:*:*:*:*",
			"software", "any", "", "any", "", false,
		},
		{
			"dash vendor and product",
			"cpe:2.3:a:-:-:1.0:*:*:*:*:*:*:*",
			"software", "any", "", "10", "any::10", true,
		},
	}

	for _, c := range cases {

		product := ToProduct(c.value)

		if product.Type != c.wantType {
			t.Errorf("%s: ToProduct() Type = %q, expected %q", c.name, product.Type, c.wantType)
		}

		if product.Vendor != c.wantVendor {
			t.Errorf("%s: ToProduct() Vendor = %q, expected %q", c.name, product.Vendor, c.wantVendor)
		}

		if product.Product != c.wantProduct {
			t.Errorf("%s: ToProduct() Product = %q, expected %q", c.name, product.Product, c.wantProduct)
		}

		if product.Version != c.wantVersion {
			t.Errorf("%s: ToProduct() Version = %q, expected %q", c.name, product.Version, c.wantVersion)
		}

		if product.Name != c.wantName {
			t.Errorf("%s: ToProduct() Name = %q, expected %q", c.name, product.Name, c.wantName)
		}

		if product.IsValid() != c.wantValid {
			t.Errorf("%s: ToProduct() IsValid() = %v, expected %v", c.name, product.IsValid(), c.wantValid)
		}

	}

}

func TestProductToProductLegacy(t *testing.T) {

	cases := []struct {
		name        string
		value       string
		wantType    string
		wantVendor  string
		wantProduct string
		wantVersion string
		wantName    string
		wantValid   bool
	}{
		{
			"application with update",
			"cpe:/a:vendor:product:1.0:beta",
			"software", "vendor", "product", "10-beta", "vendor:product:10-beta", true,
		},
		{
			"hardware with version",
			"cpe:/h:vendor:product:2.0",
			"hardware", "vendor", "product", "20", "vendor:product:20", true,
		},
		{
			"operating system",
			"cpe:/o:vendor:product",
			"system", "vendor", "product", "any", "vendor:product:any", true,
		},
		{
			"application",
			"cpe:/a:vendor:product",
			"software", "vendor", "product", "any", "vendor:product:any", true,
		},
		{
			"architecture suffix",
			"cpe:/a:vendor:product:1.0_x86_64",
			"software", "vendor", "product", "10", "vendor:product:10", true,
		},
		{
			"not a cpe",
			"not-a-cpe",
			"", "", "", "", "", false,
		},
	}

	for _, c := range cases {

		product := ToProduct(c.value)

		if product.Type != c.wantType {
			t.Errorf("%s: ToProduct() Type = %q, expected %q", c.name, product.Type, c.wantType)
		}

		if product.Vendor != c.wantVendor {
			t.Errorf("%s: ToProduct() Vendor = %q, expected %q", c.name, product.Vendor, c.wantVendor)
		}

		if product.Product != c.wantProduct {
			t.Errorf("%s: ToProduct() Product = %q, expected %q", c.name, product.Product, c.wantProduct)
		}

		if product.Version != c.wantVersion {
			t.Errorf("%s: ToProduct() Version = %q, expected %q", c.name, product.Version, c.wantVersion)
		}

		if product.Name != c.wantName {
			t.Errorf("%s: ToProduct() Name = %q, expected %q", c.name, product.Name, c.wantName)
		}

		if product.IsValid() != c.wantValid {
			t.Errorf("%s: ToProduct() IsValid() = %v, expected %v", c.name, product.IsValid(), c.wantValid)
		}

	}

}

func TestProductIsValid(t *testing.T) {

	valid := func() Product {

		product := NewProduct("software")
		product.SetVendor("acme")
		product.SetProduct("widget")
		product.SetVersion("1.0")
		product.SetState("published")

		return product

	}

	cases := []struct {
		name    string
		product Product
		want    bool
	}{
		{"default", Product{}, false},
		{"new product", NewProduct("software"), false},
		{"valid", valid(), true},
		{"missing name", func() Product { value := valid(); value.Name = ""; return value }(), false},
		{"invalid state", func() Product { value := valid(); value.SetState("invalid"); return value }(), false},
		{"missing vendor", func() Product { value := valid(); value.Vendor = ""; return value }(), false},
		{"missing type", func() Product { value := valid(); value.Type = ""; return value }(), false},
	}

	for _, c := range cases {

		if c.product.IsValid() != c.want {
			t.Errorf("%s: IsValid() = %v, expected %v", c.name, c.product.IsValid(), c.want)
		}

	}

}

func TestProductEqual(t *testing.T) {

	base := NewProduct("software")
	base.SetVendor("acme")
	base.SetProduct("widget")
	base.SetVersion("1.0")
	base.SetState("published")

	if base.Equal(base) == false {
		t.Fatal("a product should equal itself")
	}

	cases := []struct {
		name  string
		value Product
	}{
		{"name", func() Product { value := base; value.Name = "other"; return value }()},
		{"version", func() Product { value := base; value.Version = "20"; return value }()},
		{"vendor", func() Product { value := base; value.Vendor = "other"; return value }()},
		{"type", func() Product { value := base; value.Type = "hardware"; return value }()},
	}

	for _, c := range cases {

		if base.Equal(c.value) == true {
			t.Errorf("Equal() returned true for a different %s", c.name)
		}

	}

	ignored := base
	ignored.Product = "other"
	ignored.State = "edited"

	if base.Equal(ignored) == false {
		t.Fatal("Equal() should ignore the product and state fields")
	}

}

func TestProductHash(t *testing.T) {

	base := NewProduct("software")
	base.SetVendor("acme")
	base.SetProduct("widget")
	base.SetVersion("1.0")
	base.SetState("published")

	hash := base.Hash()

	if hash == "" {
		t.Fatal("Hash() returned an empty string for a populated product")
	}

	if base.Hash() != hash {
		t.Fatal("Hash() is not deterministic")
	}

	empty := Product{}

	if empty.Hash() != "" {
		t.Fatal("Hash() should be empty without a name")
	}

	cases := []struct {
		name  string
		value Product
	}{
		{"name", func() Product { value := base; value.Name = "other"; return value }()},
		{"version", func() Product { value := base; value.Version = "20"; return value }()},
		{"vendor", func() Product { value := base; value.Vendor = "other"; return value }()},
		{"type", func() Product { value := base; value.Type = "hardware"; return value }()},
	}

	for _, c := range cases {

		if c.value.Hash() == hash {
			t.Errorf("Hash() did not change for a different %s", c.name)
		}

	}

	ignored := base
	ignored.Product = "other"
	ignored.State = "edited"

	if ignored.Hash() != hash {
		t.Fatal("Hash() should ignore the product and state fields")
	}

}

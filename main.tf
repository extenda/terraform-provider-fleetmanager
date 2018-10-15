provider "fleetmanager" {
  host = "tic2d7to6f.execute-api.eu-west-1.amazonaws.com"
  tenant_id = "???"
}

resource "fleetmanager_brand" "nicks" {
  name = "Nick's Custom Brand"
}

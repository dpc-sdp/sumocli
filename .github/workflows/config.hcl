source = ["./sumocli"]
bundle_id = "com.thepublicclouds.sumocli"

apple_id {
  username = "@env:AC_USERNAME"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "08C9A4B04024CDEF2C5D93593710C60591FB4614"
}

dmg {
  output_path = "sumocli.dmg"
  volume_name = "sumocli"
}

zip {
  output_path = "sumocli.zip"
}

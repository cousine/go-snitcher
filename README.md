Go Snitcher
===========

An unofficial go client for [Dead Man's Snitch](https://deadmanssnitch.com/).

Usage
=====

Using go-snitcher is dead simple:

    import(
      "github.com/cousine/go-snitcher"
    )

    func main() {
      snitcher := snitcher.NewSnitch("YOUR_TOKEN_HERE")

      // normal checkin
      err := snitcher.Snitch()
      if err != nil {
        panic(err.Error())
      }

      // message checkin
      err = snitcher.SnitchWithMessage("This is a message checkin")
      if err != nil {
        panic(err.Error())
      }
    }

Copyright
=========

Copyright (c) 2015 Omar Mekky. See LICENSE.txt for further details.



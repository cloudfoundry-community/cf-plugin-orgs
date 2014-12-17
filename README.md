cf-plugin-orgs
==============

Proof of concept of a plugin for the [Cloudfoundry CLI](https://github.com/cloudfoundry/cli)
Use "curl" command to call api endpoint 


# Usage
**Compile plugin**
```
#Godep or #go get up to you
go build
```
**Install plugin**
```
cf install-plugin cf-plugin-orgs 
Installing plugin cf-plugin-orgs...
OK
Plugin rpaasPlugin successfully installed.
```
**Use the plugin**
```
#Connect as admin
cf login ...
#Use the plugin command
cf rorgs
OR
cf rpaas-orgs
```

Contributing
============

Pull Requests
---------------------

Pull Requests should be made against the `develop` branch.



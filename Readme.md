# pm
A command line password manager
## Usage:
pm uses symetric encryption to keep your data secure.
On first usage, pm will ask you which device you want to use, but whatever you choose
will be permanently bound as the hardward key for your data. If you decide to use a usb drive, 
your data on the drive will *NOT* be affected in any way, and you can continue to use 
your device as normal. All your files are serialized and encrypted until a pm shell is lauched with 
the `pm` command.

## `pm` Shell Commands

| CMD   |   Argument   |                Description                       |
|-------|--------------|--------------------------------------------------|
|  get  | service-name | Gets user credentials for a service              |
|  upd  | service-name | Updates credentials for an existing service      |
|  add  | service-name | Add a new service's credentials to pm            |
|  del  | service-name | Permanently delete a services credentials        |
|  ls   |     N/A      | List all services (does not display credentials) |

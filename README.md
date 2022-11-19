# Nine-headed-Bird
A high performance, parallel remote ssh shell


Nine-headed-Bird is a ssh terminal that allows you to execute a command over ssh on multiple hosts with one command.

OPTIONS

``` bash
./Nine-headed-Bird -help
Usage of ./Nine-headed-Bird:
  -O int
    	host port (default 22)
  -c string
    	command
  -d	debug
  -i string
    	ipaddress
    		example:
    		"192.168.0.1 192.168.0.1"
    		"root:123456@192.168.0.1:23 admin@192.168.0.1 192.168.0.1" 
  -n uint
    	parallel number (default 64)
  -p string
    	password
  -t int
    	connect timeout (default 5)
  -u string
    	user (default "root")
```

EXAMPLES

#### get hostname to 3 servers

``` bash
        $ ./Nine-headed-Bird  -c "hostname" -p 123456 -i "192.168.0.1 192.168.0.2 192.168.0.3"
```

outputs

``` json
{
	"Tasks": [
		{
			"Host": "root:123456@192.168.0.1:22",
			"Elapsed": "177.712039ms",
			"Index": 1,
			"Result": [
				"Ansible",
				""
			],
			"Error": ""
		},
		{
			"Host": "root:123456@192.168.0.3:22",
			"Elapsed": "181.017892ms",
			"Index": 2,
			"Result": [
				"Ansible",
				""
			],
			"Error": ""
		},
		{
			"Host": "root:123456@192.168.0.2:22",
			"Elapsed": "186.795931ms",
			"Index": 0,
			"Result": [
				"Ansible",
				""
			],
			"Error": ""
		}
	],
	"Count": 3,
	"Success": true,
	"Elapsed": "187.054418ms"
}

```

#### get hostnames of 3 servers, with specify parameters respectively

``` bash
        $ ./Nine-headed-Bird  -c "hostname" -p 123456 -i "root:654321@192.168.0.1:23 192.168.0.2 192.168.0.3"
```

outputs

``` json
{
	"Tasks": [
		{
			"Host": "root:654321@192.168.0.1:23",
			"Elapsed": "177.712039ms",
			"Index": 1,
			"Result": [
				"Ansible",
				""
			],
			"Error": ""
		},
		{
			"Host": "root:123456@192.168.0.3:22",
			"Elapsed": "181.017892ms",
			"Index": 2,
			"Result": [
				"Ansible",
				""
			],
			"Error": ""
		},
		{
			"Host": "root:123456@192.168.0.2:22",
			"Elapsed": "186.795931ms",
			"Index": 0,
			"Result": [
				"Ansible",
				""
			],
			"Error": ""
		}
	],
	"Count": 3,
	"Success": true,
	"Elapsed": "187.054418ms"
}

```


#### get hostnames of 254 servers, with 128 parallel routine,and set the connection timeout to 1 second

``` bash
        $ ./Nine-headed-Bird -t 1 -n 128 -c "hostname" -p 123456 -i "``seq 1 254|xargs -I[] echo '192.168.0.'[]|xargs``"
```

outputs

``` json
{
    "Tasks": [
        {
            "Host": "root:123456@192.168.0.1:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.2:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.3:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.4:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.5:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.6:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.7:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.8:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.9:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.10:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.11:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.12:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.13:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.14:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.15:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.16:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.17:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.18:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.19:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.20:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.21:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.22:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.23:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.24:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.25:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.26:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.27:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.28:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.29:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.30:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.31:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.32:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.33:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.34:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.35:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.36:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.37:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.38:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.39:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.40:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.41:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.42:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.43:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.44:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.45:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.46:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.47:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.48:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.49:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.50:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.51:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.52:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.53:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.54:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.55:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.56:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.57:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.58:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.59:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.60:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.61:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.62:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.63:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.64:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.65:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.66:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.67:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.68:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.69:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.70:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.71:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.72:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.73:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.74:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.75:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.76:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.77:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.78:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.79:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.80:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.81:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.82:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.83:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.84:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.85:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.86:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.87:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.88:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.89:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.90:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.91:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.92:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.93:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.94:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.95:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.96:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.97:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.98:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.99:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.100:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.101:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.102:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.103:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.104:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.105:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.106:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.107:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.108:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.109:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.110:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.111:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.112:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.113:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.114:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.115:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.116:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.117:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.118:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.119:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.120:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.121:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.122:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.123:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.124:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.125:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.126:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.127:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.128:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.129:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.130:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.131:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.132:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.133:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.134:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.135:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.136:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.137:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.138:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.139:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.140:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.141:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.142:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.143:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.144:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.145:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.146:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.147:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.148:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.149:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.150:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.151:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.152:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.153:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.154:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.155:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.156:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.157:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.158:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.159:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.160:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.161:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.162:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.163:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.164:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.165:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.166:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.167:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.168:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.169:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.170:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.171:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.172:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.173:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.174:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.175:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.176:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.177:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.178:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.179:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.180:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.181:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.182:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.183:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.184:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.185:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.186:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.187:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.188:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.189:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.190:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.191:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.192:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.193:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.194:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.195:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.196:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.197:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.198:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.199:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.200:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.201:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.202:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.203:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.204:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.205:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.206:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.207:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.208:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.209:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.210:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.211:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.212:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.213:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.214:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.215:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.216:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.217:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.218:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.219:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.220:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.221:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.222:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.223:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.224:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.225:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.226:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.227:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.228:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.229:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.230:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.231:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.232:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.233:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.234:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.235:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.236:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.237:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.238:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.239:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.240:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.241:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.242:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.243:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.244:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.245:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.246:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.247:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.248:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.249:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.250:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.251:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.252:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.253:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        },
        {
            "Host": "root:123456@192.168.0.254:21",
            "Elapsed": "177.712039ms",
            "Index": 1,
            "Result": [
                "Ansible",
                ""
            ],
            "Error": ""
        }
    ],
    "Count": 254,
    "Success": true,
    "Elapsed": "1.008980628s"
}

```




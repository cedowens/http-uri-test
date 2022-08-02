# http-uri-test

Simple client and server in go to test data exfiltration over http uri paths. Steps:

## Server Side:
1. > cd server

2. > go build

3. > ./server

This will start your golang http server that will listen for and parse inbound http request uri values. Reminder: place the appropriate firewall rules in front of this server to ensure that only your testing client can access port 80 on this server


## Client Side:
1. > ulimit -n 1000000

1. > cd client

2. Replace the placeholder string "domain-or-IP-here" in the client.go file with the domain or IP address that you are using for your web server.

3. > go build

4. > ./client <path_to_file_to send>

This will take whatever file you point it to, break it into 200 character chunks of hex encoded text, and send it to the server in the form of **GET /[200_char_hex_encoded_chunk** to the target server. The target server will in turn parse these uri values, hex decode them, and write them to a file in the current directory named "outfile"

This makes for a good test to see if files exfiltrated in this manner are detected in your environment.

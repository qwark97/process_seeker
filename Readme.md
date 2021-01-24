Simple service which runs HTTP server with one, GET endpoint to search running processes

route: "/search"<br>
method: GET<br>
body:<br>
```json
{"processName": "full process name"}
```

As a result response will be:
1) If process has been found running:
```json
{"status":0,"message":"Process exists"}
```
2) If process has not been found:
```json
{"status":1,"message":"Process does not exist"}
```
3) If some error occured:
```json
{"status":-1,"message":"Some error message"}
```
<br>
<br>
<br>
To run it needs one parameter which is port value.<br>

e.g:
```bash
./process_seeker 8080
```

It needs two environmental variables to work:<br>
PROCESS_SEEKER_PYTHON_INTERPRETER<br>
PROCESS_SEEKER_SCRIPT_PATH
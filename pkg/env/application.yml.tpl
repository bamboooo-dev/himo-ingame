himoMySQL:
  host: "{{ .HIMO_DB_HOST }}"
  port: {{ .HIMO_DB_PORT }}
  database: "{{ .HIMO_DB_DATABASE }}"
  user: "{{ .HIMO_DB_USER }}"
  password: "{{ .HIMO_DB_PASSWORD }}"
  parseTime: {{ .HIMO_DB_PARSE_TIME }}

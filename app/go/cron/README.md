Este ejemplo muestra como configurar un cron basado en Temporal en un sistema distribuido.

Pasos para ejecutar el ejemplo:
1) Servicio de temporal ejecutandose. 
2) Ejecutar el worker
```
go run cron/worker/main.go 
```

3) Ejecutar el starter para que de alli corra el proceso de cron. A partir de ahora de acuerdo a la configuracion el workflow se ejecutara cada minuto.

```
go run cron/starter/main.go
```

El starter contiene la linea (ver abajo) que configura el cron:
CronSchedule: "* * * * *"


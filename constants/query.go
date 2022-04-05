package constants

const ALTITUDE_GET = "SELECT id, valore, data_inserimento FROM altitudine order by id desc limit 100"
const ALTITUDE_GET_LAST = "SELECT id, valore, data_inserimento FROM altitudine where id = (select max(id) from altitudine)"
const ALTITUDE_GET_LASTHOUR = "SELECT id,valore,data_inserimento FROM altitudine where data_inserimento  >= '%s' AND data_inserimento <= '%s'"
const ALTITUDE_GET_SHOWDATA = "WITH t AS (SELECT id,valore,data_inserimento FROM altitudine ORDER BY data_inserimento DESC LIMIT %d) SELECT id,valore,data_inserimento FROM t ORDER BY data_inserimento ASC"
const ALTITUDE_POST_DATA = "insert into altitudine (valore,data_inserimento) values (%.2f,CURRENT_TIMESTAMP) RETURNING id"

const ACTIVITY_ISACTIVE = "select count(id) as contatore from attivita where  data_inserimento <=now() and data_inserimento >= now() - INTERVAL '1 MINUTES'"
const ACTIVITY_GET = "SELECT id,0, data_inserimento FROM attivita order by id desc limit 100"
const ACTIVITY_GET_LAST = "SELECT id, 0, data_inserimento FROM attivita where id = (select max(id) from attivita)"
const ACTIVITY_GET_LASTHOUR = "SELECT id,0, data_inserimento FROM attivita where data_inserimento  >= '%s' AND data_inserimento <= '%s'"

const MESSAGE_GET_LAST = "SELECT id, messaggio, data_inserimento FROM messaggio where id = (select max(id) from attivita)"
const MESSAGE_GET_LASTHOUR = "SELECT id,messaggio, data_inserimento FROM messaggio where data_inserimento  >= '%s' AND data_inserimento <= '%s'"
const MESSAGE_POST_DATA = "insert into messaggio (messaggio,data_inserimento) values (%s,CURRENT_TIMESTAMP) RETURNING id"

const RAIN_GET = "SELECT id, valore, data_inserimento FROM pioggia order by id desc limit 100"
const RAIN_GET_LAST = "SELECT id, valore, data_inserimento FROM pioggia where id = (select max(id) from pioggia)"
const RAIN_GET_LAST_HOUR = "SELECT id,valore,data_inserimento FROM pioggia where data_inserimento  >= '%s' AND data_inserimento <= '%s'"
const RAIN_POST_DATA = "insert into pioggia (valore,data_inserimento) values (%d,CURRENT_TIMESTAMP) RETURNING id"
const RAIN_GET_SHOWDATA = "WITH t AS (SELECT id,valore,data_inserimento FROM pioggia ORDER BY data_inserimento DESC LIMIT %d) SELECT id,valore,data_inserimento FROM t ORDER BY data_inserimento ASC"

const PRESSURE_GET = "SELECT id, valore, data_inserimento FROM pressione order by id desc limit 100"
const PRESSURE_GET_LAST = "SELECT id, valore, data_inserimento FROM pressione where id = (select max(id) from pressione)"
const PRESSURE_GET_LAST_HOUR = "SELECT id,valore,data_inserimento FROM pressione where data_inserimento  >= '%s' AND data_inserimento <= '%s'"
const PRESSURE_POST_DATA = "insert into pressione (valore,data_inserimento) values (%.2f,CURRENT_TIMESTAMP) RETURNING id"
const PRESSURE_GET_SHOWDATA = "WITH t AS (SELECT id,valore,data_inserimento FROM pressione ORDER BY data_inserimento DESC LIMIT %d) SELECT id,valore,data_inserimento FROM t ORDER BY data_inserimento ASC"

const TEMPERATURE_GET = "SELECT id, valore, data_inserimento FROM temperatura order by id desc limit 100"
const TEMPERATURE_GET_LAST = "SELECT id, valore, data_inserimento FROM temperatura where id = (select max(id) from temperatura)"
const TEMPERATURE_GET_LAST_HOUR = "SELECT id,valore,data_inserimento FROM temperatura where data_inserimento  >= '%s' AND data_inserimento <= '%s'"
const TEMPERATURE_POST_DATA = "insert into temperatura (valore,data_inserimento) values (%.2f,CURRENT_TIMESTAMP) RETURNING id"
const TEMPERATURE_GET_SHOWDATA = "WITH t AS (SELECT id,valore,data_inserimento FROM temperatura ORDER BY data_inserimento DESC LIMIT %d) SELECT id,valore,data_inserimento FROM t ORDER BY data_inserimento ASC"

const HUMIDITY_GET = "SELECT id, valore, data_inserimento FROM umidita order by id desc limit 100"
const HUMIDITY_GET_LAST = "SELECT id, valore, data_inserimento FROM umidita where id = (select max(id) from umidita)"
const HUMIDITY_GET_LAST_HOUR = "SELECT id,valore,data_inserimento FROM umidita where data_inserimento  >= '%s' AND data_inserimento <= '%s'"
const HUMIDITY_POST_DATA = "insert into umidita (valore,data_inserimento) values (%.2f,CURRENT_TIMESTAMP) RETURNING id"
const HUMIDITY_GET_SHOWDATA = "WITH t AS (SELECT id,valore,data_inserimento FROM umidita ORDER BY data_inserimento DESC LIMIT %d) SELECT id,valore,data_inserimento FROM t ORDER BY data_inserimento ASC"

const PREVISION_GET = "select 'TIPO PRESSIONE' as Dato, case when round(valore)>1013 then 'ALTA PRESSIONE' else 'BASSA PRESSIONE' end as Valore from  pressione where  id = (select max(id) from pressione) union select 'PRESSIONE MINIMA' as Dato, cast(min(round(valore)) as varchar) as Valore from   (select * from pressione where  data_inserimento <=now() and data_inserimento >= now() - INTERVAL '3 HOURS') t union select 'PRESSIONE MASSIMA' as Dato, cast(max(round(valore)) as varchar) as Valore from  (select * from pressione where  data_inserimento <=now() and data_inserimento >= now() - INTERVAL '3 HOURS') t union select 'TEMPERATURA MINIMA' as Dato, cast(min(round(valore)) as varchar) as Valore from  (select * from temperatura where  data_inserimento <=now() and data_inserimento >= now() - INTERVAL '3 HOURS') t union select 'TEMPERATURA MASSIMA' as Dato,cast(max(round(valore)) as varchar) as Valore from (select * from temperatura where  data_inserimento <=now() and data_inserimento >= now() - INTERVAL '3 HOURS') t"

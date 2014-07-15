=======
Prueba
=======
El código de este repositorio git es un código escrito en Go que obtiene una configuración de firewall mediante una llamada a la API de servidores de Tapp. Se trata de una llamada autenticada tanto para cliente como para servidor por SSL. La prueba propuesta es tomando como base este código escribir en Go un programa que obtenga un listado de scripts a ejecutar, los ejecute y comunique los datos de la ejecución a la API.
	1. Para obtener de la API el listado de scripts (script_characterizations) de arranque, el procedimiento sería similar al mostrado en el código de base, pero cambiando el path a '/blueprint/script_characterizations?type=boot'. La respuesta a esta petición sería un array en json como el siguiente:
		[{"execution_order":1,"parameter_values":{},"type":"boot","uuid":"53c3b86e63051f336b00036f","script":{"code":"#!/bin/bash\nhostname\npwd","uuid":"53c3b79963051f7a8800036b","attachment_paths":[]}},{"execution_order":4611686018427387903,"parameter_values":{},"type":"boot","uuid":"53c3b73163051f6d46000367","script":{"code":"echo .","uuid":"4ea19048d907b10559000003","attachment_paths":[]}}]
	2. De la ejecución del código de estos scripts se debería retener el output del script, el código de salida y los timestamps de inicio y finalización de la ejecución. Se sugiere volcar el código de cada script a un fichero y ejecutarlo con sh.
	3. Para comunicar mediante la API de servidores de Tapp los datos de la ejecución de cada script se debe hacer un POST al path '/blueprint/script_conclusions', con Content-Type: application/json, de los datos retenidos y el uuid del script_characterization. Esto es, el payload debería ser similar a
		{"script_conclusion":{"script_characterization_id":"53c3b86e63051f336b00036f","output":"iMac.local\n/Users/pbanos/Code/gocode/src/github.com/jpgriffo/tapp-client\n","exit_code":0,"started_at":"2014-07-14T13:37:21.283363+02:00","finished_at":"2014-07-14T13:37:21.323643+02:00"}}
	La API de servidores de Tapp actualmente responde con código 500 en lugar de un código 400 si alguno de los campos esperados no se proporciona. En caso de éxito devuelve una respuesta con código 201.

==========
Resolución
==========
Clonar este repositorio en vuestro equipo y trabajad sobre el para resolver la prueba. Los ficheros correspondientes al certificado SSL de cliente y a la CA, configurados en tapp/client.xml, os los proporcionaremos por correo: por favor, absteneros de subirlos al repositorio. Cuando hayais terminado, comunicadnoslo por email indicando la URL del repositorio. 


=======================
Criterios de valoración
=======================
* Valoraremos los apartados independientemente, si os 'atrancáis' con un apartado no pasa nada, pasad al siguiente y mandadnos lo que tengais
* Como bonus valoraríamos tests del código desarrollado, o incluso mejor, la aplicación de TDD/BDD
* No solo valoramos el resultado final si no también el proceso de desarrollo: examinaremos el historial de git

La realización de la prueba no es obligatoria para permanecer en el proceso de selección, pero si que ayuda bastante puesto que nos permite evaluaros en base a algo más que el CV. 

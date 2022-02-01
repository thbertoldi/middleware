<script>
 import { onMount } from 'svelte';
 let appConfig = {"http": {"port": 0}, "mqtt": {"host": "", "port": 0, "topic": [], "id": ""}};
 let allowJoin = true;
 async function sendConfig() {
		 const res = await fetch("http://localhost:5000/api/config", {
			   method: 'POST',
			   body: JSON.stringify(appConfig)
		 })
 };
 onMount(async () => {
     console.log(location.href + "api/config");
     fetch(location.href + "api/config")
         .then(response => response.json())
         .then(data => {
		         console.log(data);
             appConfig = data
         }).catch(error => {
             console.log(error);
             return [];
         });
 });
</script>

<h1 class="bg-primary text-light text-center">IoT Middleware</h1>

<form>
    <h3>
        Aplicação Web
    </h3>
    <div class="mb-3">
        <label for="httpPort" class="form-label">Porta para o Servidor</label>
        <input bind:value={appConfig.http.port} type="text" class="form-control" id="httpPort" aria-describedby="portHelp">
        <div id="portHelp" class="form-text">Porta TCP para execução do servidor Web</div>
    </div>
    <hr>
    <h3>
        Cliente MQTT
    </h3>
    <div class="mb-3">
        <label for="mqttPort" class="form-label">Porta para o Cliente MQTT</label>
        <input bind:value={appConfig.mqtt.port} type="number" class="form-control" id="mqttPort" aria-describedby="mqPortHelp">
        <div id="mqPortHelp" class="form-text">Porta TCP para conexão do cliente MQTT com o Broker</div>
    </div>
    <div class="mb-3">
        <label for="mqttHost" class="form-label">Endereço para o Cliente MQTT</label>
        <input bind:value={appConfig.mqtt.host} type="text" class="form-control" id="mqttHost" aria-describedby="mqHostHelp">
        <div id="mqHostHelp" class="form-text">Endereço do Broker para conexão</div>
    </div>
    <div class="mb-3">
        <label for="mqttId" class="form-label">Identificador para o Cliente MQTT</label>
        <input bind:value={appConfig.mqtt.id} type="text" class="form-control" id="mqttPort" aria-describedby="mqIdHelp">
        <div id="mqIdHelp" class="form-text">Identificador para o cliente MQTT</div>
    </div>
    {#each appConfig.mqtt.topic as t}
        <label for="mqttTopic" class="form-label">Tópico</label>
        <input bind:value={t} type="text" class="form-control" id="mqttTopic">
    {/each}
    <hr>
    <h3>
        Gerais
    </h3>
    <div class="mb-3 form-check">
        <input bind:value={allowJoin} type="checkbox" class="form-check-input" id="checkJoin">
        <label class="form-check-label" for="checkJoin">Permiter auto-cadastro de dispositivos</label>
    </div>
    <button type="submit" class="btn btn-primary" on:click="{sendConfig}">Enviar</button>
</form>

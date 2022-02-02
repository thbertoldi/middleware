<script>
 import { onMount } from 'svelte';
 let appConfig = {"http": {"port": 0}, "mqtt": {"host": "", "port": 0, "topic": [], "id": ""}};
 let allowJoin = true;
 async function sendConfig() {
     const res = await fetch(location.origin + "/api/config", {
         method: 'POST',
         body: JSON.stringify(appConfig)
     })
 };
 onMount(async () => {
     fetch(location.origin + "/api/config")
         .then(response => response.json())
         .then(data => {
             appConfig = data
         }).catch(error => {
             return [];
         });
 });
</script>

<html lang="en">
    <body class="bg-light">
        <h1 class="bg-secondary text-light text-center">IoT Middleware</h1>
        <div class="container">
            <form class="row mx-auto offset-lg-2">
                <div class="row mb-3">
                    <label for="httpPort" class="col-sm-2 col-form-label">Porta para o Servidor</label>
                    <div class="col-sm-4">
                        <input bind:value={appConfig.http.port} type="number" class="form-control" id="httpPort">
                    </div>
                </div>
                <div class="row mb-3">
                    <label for="mqttPort" class="col-sm-2 col-form-label">Porta para o Cliente MQTT</label>
                    <div class="col-sm-4">
                        <input bind:value={appConfig.mqtt.port} type="number" class="form-control" id="mqttPort">
                    </div>
                </div>
                <div class="row mb-3">
                    <label for="mqttHost" class="col-sm-2 col-form-label">Endereço para o Cliente MQTT</label>
                    <div class="col-sm-4">
                        <input bind:value={appConfig.mqtt.host} type="text" class="form-control" id="mqttHost">
                    </div>
                </div>
                <div class="row mb-3">
                    <label for="mqttId" class="col-sm-2 col-form-label">Identificador para o Cliente MQTT</label>
                    <div class="col-sm-4">
                        <input bind:value={appConfig.mqtt.id} type="text" class="form-control" id="mqttPort">
                    </div>
                </div>
                {#each appConfig.mqtt.topic as t}
                    <div class="row mb-3">
                        <label for="mqttTopic" class="col-sm-2 col-form-label">Tópico</label>
                        <div class="col-sm-4">
                            <input bind:value={t} type="text" class="form-control" id="mqttTopic">
                        </div>
                    </div>
                {/each}
                <div class="row mb-3">
                    <label class="col-sm-2 col-form-label">Dispositivos</label>
                    <div class="col-sm-1">
                        <input bind:value={allowJoin} type="checkbox" class="form-check-input" id="checkJoin">
                    </div>
                    <label class="col-sm-4 form-check-label" for="checkJoin">Permitir auto-cadastro de dispositivos</label>
                </div>
                <div class="row mb-3">
                    <div class="col-sm-6">
                        <button type="submit" class="btn btn-primary" on:click="{sendConfig}">Enviar</button>
                    </div>
                </div>
            </form>
        </div>
    </body>
</html>

<template>
    <div class="q-pa-md">
        <q-card flat bordered>
            <q-card-section class="row justify-center">
                <div class="q-pa-md" style="width: 100%">
                    <ImportarComponent />
                </div>
            </q-card-section>
        </q-card>
    </div>
</template>
  
<script>
    import ImportarComponent from "components/importacao/ImportarComponent.vue"

    export default {
        components: {
            ImportarComponent,
        },
        setup() {
            return {
            };
        },
    };
</script>
  
import API from "@/http/index"

export default {

    listar(){
        return API.get('/classification-segment')
    },
}
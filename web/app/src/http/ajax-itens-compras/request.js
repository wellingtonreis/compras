// import { http } from "./config";
import API from "@/http/index"

export default {

    listar(cotacao){
        return API.get(`/purchase-items/${cotacao}`)
    },

    atualizar(cotacao, item){
        return API.put(`/purchase-items/${cotacao}`, item)
    }
}
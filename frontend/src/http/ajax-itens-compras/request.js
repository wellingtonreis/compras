// import { http } from "./config";
import API from "@/http/index"

export default {

    listar(cotacao){
        return API.get(`/purchase-items/${cotacao}/get`)
    },

    atualizar(cotacao, item){
        return API.put(`/purchase-items/${cotacao}/put`, item)
    },

    remover(cotacao, item){
        return API.put(`/purchase-items/${cotacao}/delete`, item)
    }
}
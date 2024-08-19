// import { http } from "./config";
import API from "@/http/index"

export default {

    atualizar(historicoCotacao, cotacaoId){
        return API.put(`/quotation-history/classification-segment/${cotacaoId}/put`, historicoCotacao)
    },

    listar(filter){
        return API.post('/quotation-history', filter)
    },
}
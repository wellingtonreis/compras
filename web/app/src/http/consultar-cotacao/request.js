// import { http } from "./config";
import API from "@/http/index"

export default {

    // cadastrar(calendario){
    //     return API.post('multicarteira/calendario-atividade/salvar', calendario)
    // },

    listData(filter){
        return API.post('/quotation-history', filter)
    },

    // excluir(id){
    //     return API.delete('multicarteira/calendario-atividade/excluir/'+id)
    // }
}
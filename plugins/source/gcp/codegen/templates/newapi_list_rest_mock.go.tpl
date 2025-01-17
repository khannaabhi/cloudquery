// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	{{if .ProtobufImport}}
  pb "{{.ProtobufImport}}"
  {{end}}
)


func create{{.SubService | ToCamel}}(mux *httprouter.Router) error {  
  var item pb.{{.ResponseStructName}}
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	emptyStr := ""
	item.NextPageToken = &emptyStr
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func Test{{.SubService | ToCamel}}(t *testing.T) {
	client.MockTestRestHelper(t, {{.SubService | ToCamel}}(), create{{.SubService | ToCamel}}, client.TestOptions{})
}

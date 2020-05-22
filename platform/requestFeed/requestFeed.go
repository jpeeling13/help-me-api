package requestFeed

type Request struct {
	Description     string `json: "description"`
	RequesterName string `json: "requesterName"`
}

type TempDataStore struct {
	Requests []Request
}

func New() *TempDataStore {
	return &TempDataStore{
		[]Request{},
	}
}

func (tds *TempDataStore) Add(request Request){
	tds.Requests = append(tds.Requests, request)
}

func (tds *TempDataStore) GetAll() []Request {
	return tds.Requests
}
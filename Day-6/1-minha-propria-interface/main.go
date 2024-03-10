package main

type DataBucket struct {
	Data []byte
}

type DataLoader interface {
	LoadData(filename string, bucket DataBucket) error
}

// XLS
type XLSDataLoader struct{}

func (xlsdataloader *XLSDataLoader) LoadData(filename string, bucket DataBucket) error {
	// Carrega dados de um arquivo de nome filename dentro desse databucket, fazendo as transformações de dados necessarias para um XLS
	return nil
}

// CSV
type CSVDataLoader struct{}

func (csvdataloader *CSVDataLoader) LoadData(filename string, bucket DataBucket) error {
	// Carrega dados de um arquivo de nome filename dentro desse databucket, fazendo as transformações de dados necessarias para um CSV
	return nil
}

func LoadData(filename string, dataloader DataLoader) DataBucket {
	bucket := DataBucket{}
	dataloader.LoadData(filename, bucket)
	return bucket
}

func main() {
	LoadData("arquivo.csv", &CSVDataLoader{})
	LoadData("arquivo.xls", &XLSDataLoader{})
}

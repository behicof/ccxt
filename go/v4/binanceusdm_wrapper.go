package ccxt

type Binanceusdm struct {
   *binanceusdm
   Core *binanceusdm
}

func NewBinanceusdm(userConfig map[string]interface{}) Binanceusdm {
   p := &binanceusdm{}
   p.Init(userConfig)
   return Binanceusdm{
       binanceusdm: p,
       Core:  p,
   }
}

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code


func (this *Binanceusdm) TransferIn(code string, amount interface{}, options ...TransferInOptions) (TransferEntry, error) {

    opts := TransferInOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.TransferIn(code, amount, params)
    if IsError(res) {
        return TransferEntry{}, CreateReturnError(res)
    }
    return NewTransferEntry(res), nil
}
func (this *Binanceusdm) TransferOut(code string, amount interface{}, options ...TransferOutOptions) (TransferEntry, error) {

    opts := TransferOutOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.TransferOut(code, amount, params)
    if IsError(res) {
        return TransferEntry{}, CreateReturnError(res)
    }
    return NewTransferEntry(res), nil
}
def GetRateResult(resRates):
    tmpBuy = []
    tmpSell = []

    status = 0

    for rate in resRates:
        if status == 0:
            tmpBuy.append(rate.text.strip())
            status = 1
            continue
        if status == 1:
            tmpSell.append(rate.text.strip())
            status = 0
            continue
    return tmpBuy, tmpSell

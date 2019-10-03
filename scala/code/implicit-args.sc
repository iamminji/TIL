def calcTax(amount: Float)(implicit rate: Float): Float = amount * rate

object SimpleStateSalesTax {
  implicit val rate: Float = 0.05F
}

case class ComplicatedSalesTaxData(baseRate: Float, isTaxHoliday: Boolean, storeId: Int)

object ComplicatedSalesTax {
  private def extraTaxRateForStore(id: Int): Float = {
    0.0F
  }

  implicit def rate(implicit cstd: ComplicatedSalesTaxData): Float =
    if (cstd.isTaxHoliday) 0.0F
    else cstd.baseRate + extraTaxRateForStroe(cstd.storeId)
}

/////////////////////////////////////////////////////////////////////////////////////////

val amount = 100F
println(s"Tax on $amount = ${calcTax(amount)}")

//////////////////////////////////////////////////////////////////////////////////////////

implicit val myStore = ComplicatedSalesTaxData(0.06F, false, 1010)
val amount = 100F
println(s"Tax on $amount = ${calcTax(amount)}")

///////////////////////////////////////////////////////////////////////////////////////////

/* ouput
Tax on 100.0 = 5.0
Tax on 100.0 = 6.0
*/


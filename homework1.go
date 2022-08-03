// İki farklı bölgede oy kullanan Seçmen Sayısı ve Aldığı oy oranına göre Seçim sonuçları
// Ekrana Toplam seçmen sayısı , Toplam alınan oy ve Aldığı oyların yüzdesini yazdıran bir program

package main

import "fmt"

func main() {

	// A ve B şehirlerinde bulunan seçmen sayısını ve A ve B şehirlerinde alınan oy sayısını tanımladık

	numberofVotersA := 20000
	numberofVotersB := 14000
	receivedVotersA := 9500
	receivedVotersB := 8000

	// toplam oy kullanan seçmen sayısını , toplam alınan oy sayısını ve toplam yüzdelik sonucu integer olarak tanımladık

	var totalnumberofVoters int
	var totalreceivedVoters int
	var totalPercentile int

	// Oy kullanan seçmen sayısını , toplam alınan oy sayısını ve yüzdelik sonucu ekrana yazdırdık

	totalnumberofVoters = calculateNumberof(numberofVotersA, numberofVotersB, receivedVotersA, receivedVotersB)
	totalreceivedVoters = calculateReceived(receivedVotersA, receivedVotersB)
	totalPercentile = calculatePercentile(totalnumberofVoters, totalreceivedVoters)

	fmt.Println("Toplam seçmen ve kullanılan oy sayısı : ", totalnumberofVoters)
	fmt.Println("Toplam alınan oy Sayısı : ", totalreceivedVoters)
	fmt.Println("Seçmen Sayısına göre alınan oyun yüzdelik toplamı : %", totalPercentile)
}

//Toplam oy kullanan seçmen sayısını tanımladık

func calculateNumberof(numberofVotersA int, numberofVotersB int, receivedVotersA int, receivedVotersB int) int {

	totalnumberofVoters := numberofVotersA + numberofVotersB

	// ****Burada A ve B 'de kullanılan toplam oy sayısının alınan oy sayısından daha az olduğu zaman sistemin hata vermesini sağladık..***

	if numberofVotersA < receivedVotersA {
		fmt.Println("A şehrinde Kullanılan oy sayısı, alınan oy sayısından küçük olamaz")

	} else if numberofVotersB < receivedVotersB {
		fmt.Println("B şehrinde kullanılan oy sayısı , alınan oy sayısından küçük olamaz")

	} else {
		fmt.Println("Seçim Sonuçları açıklanmıştır")
	}

	return totalnumberofVoters
}

//Toplam alınan oy sayısını tanımladık

func calculateReceived(receivedVotersA int, receivedVotersB int) int {
	totalreceivedVoters := receivedVotersA + receivedVotersB

	return totalreceivedVoters
}

// Toplam yüzdelik sonucu tanımladık.%50 nin altında olduğu zaman seçimi kaybetmesini , %50 ye eşit olduğu zaman seçim berabere , %50 den fazla olduğu zaman seçimi kazanmasını tanımladık..

func calculatePercentile(totalreceivedVoters int, totalnumberofVoters int) int {

	//TODO
	//totalreceivedVoters değerine 0 göndererek program çıktısını bi inceleyelim nedir sonuç ?

	totalPercentile := totalnumberofVoters * 100 / totalreceivedVoters

	if totalPercentile == 50 {
		fmt.Println("Seçim Berabere sonuçlandı.2. tura geçiniz")
	} else if totalPercentile < 50 {
		fmt.Println("Seçimi Kaybettiniz")
	} else {
		fmt.Println("Tebrikler Seçimi Kazandınız")
	}
	return totalPercentile

}

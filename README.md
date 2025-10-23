Download
bash
git clone https://github.com/sekarsister/diskrit-sifat-sifat-relasi.git
cd diskrit-sifat-sifat-relasi
Run Program
bash
go run main.go
Edit Relasi
Buka file main.go, cari bagian ini dan ganti dengan relasi Anda:

go
himpunanSaya := []Pair{
    {"1", "1"},
    {"1", "2"}, 
    {"2", "1"},
    {"2", "2"},
}
Contoh Relasi
go
// Relasi kesetaraan
{"a","a"}, {"b","b"}, {"a","b"}, {"b","a"}

// Relasi urutan  
{"1","2"}, {"2","3"}, {"1","3"}
Hasil
Program akan menampilkan analisis 6 sifat relasi:

Refleksif

Simetris

Asimetris

Anti-simetris

Transitif

Irrefleksif

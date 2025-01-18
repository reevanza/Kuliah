#include "parent.h"
#include "child.h"
#include "relasi.h"

void menu (){
    cout << endl;
    cout << "================================ MENU ================================" << endl;
    cout << "| No |                           Aksi                                |" << endl;
    cout << "----------------------------------------------------------------------" << endl;
    cout << "|  1 | Insert Penulis                                                |" << endl;
    cout << "|  2 | Insert Artikel                                                |" << endl;
    cout << "|  3 | Insert Relasi                                                 |" << endl;
    cout << "|  4 | Print Penulis                                                 |" << endl;
    cout << "|  5 | Print Artikel                                                 |" << endl;
    cout << "|  6 | Print Relasi                                                  |" << endl;
    cout << "|  7 | Delete Penulis                                                |" << endl;
    cout << "|  8 | Delete Artikel                                                |" << endl;
    cout << "|  9 | Delete Relasi                                                 |" << endl;
    cout << "| 10 | Search Penulis berdasarkan nama                               |" << endl;
    cout << "| 11 | Search Penulis berdasarkan ID                                 |" << endl;
    cout << "| 12 | Search Penulis berdasarkan Email                              |" << endl;
    cout << "| 13 | Search Artikel berdasarkan judul                              |" << endl;
    cout << "| 14 | Search Artikel berdasarkan tahun terbit                       |" << endl;
    cout << "| 15 | Search Artikel berdasarkan konten                             |" << endl;
    cout << "| 16 | Hitung Artikel berdasarkan Penulis                            |" << endl;
    cout << "| 17 | Search data artikel pada penulis tertentu                     |" << endl;
    cout << "| 18 | Exit                                                          |" << endl;
    cout << "----------------------------------------------------------------------" << endl;
    cout << "======================================================================" << endl;
    cout << endl;
    cout << "Pilih : ";
}

int main(){
    List_parent L;
    List_child C;
    List_relasi R;
    createList(L);
    createList(C);
    createList(R);
    infotype_parent x;
    infotype_child y;
    address_parent P;
    address_child Q;
    address_relasi S;

    string nama, judul;
    int choice;
    bool schoice = true;
    while (schoice) {
        menu();
        cin >> choice;
        switch (choice) {
            case 1:
                cout << "Masukkan Nama: ";
                cin >> x.nama;
                cout << "Masukkan ID: ";
                cin >> x.id;
                cout << "Masukkan Email: ";
                cin >> x.email;
                P = alokasiParent(x);

                cout << "====================" << endl;
                cout << "Pilih metode insert:" << endl;
                cout << "1. Insert First" << endl;
                cout << "2. Insert Last" << endl;
                cout << "3. Insert After" << endl;
                cout << "====================" << endl;
                cout << endl;
                cout << "Pilihan Anda: ";
                int insertChoice;
                cin >> insertChoice;

                switch (insertChoice) {
                    case 1:
                        insertFirst(L, P);
                        break;
                    case 2:
                        if (first(L) == NULL) {
                            cout << "=====================================================================" << endl;
                            cout << "List kosong, tidak dapat insert after. Data akan ditambahkan di awal." << endl;
                            cout << "=====================================================================" << endl;
                            insertFirst(L, P);
                        } else {
                            insertLast(L, P);
                        }
                        break;
                    case 3:
                        if (first(L) == NULL) {
                            cout << "=====================================================================" << endl;
                            cout << "List kosong, tidak dapat insert after. Data akan ditambahkan di awal." << endl;
                            cout << "=====================================================================" << endl;
                            insertFirst(L, P);
                        } else {
                            string targetNama;
                            printInfo(L);
                            cout << endl;
                            cout << "Masukkan nama penulis setelah siapa Anda ingin menambahkan: ";
                            cin >> targetNama;
                            address_parent Prec = findElm_nama(L, targetNama);
                            if (Prec != NULL) {
                                insertAfter(L, Prec, P);
                                cout << "Penulis berhasil ditambahkan setelah " << targetNama << "!" << endl;
                            } else {
                                cout << "Penulis dengan nama " << targetNama << " tidak ditemukan. Data tidak ditambahkan." << endl;
                                dealokasi(P);
                            }
                        }
                        break;
                    default:
                        cout << "============================================" << endl;
                        cout << "Pilihan tidak valid. Data tidak ditambahkan." << endl;
                        cout << "============================================" << endl;
                        dealokasi(P);
                        break;
                }
                break;
            case 2:
                cout << "Masukkan Judul Artikel: ";
                cin >> y.judulArtikel;
                cout << "Masukkan Tahun Terbit: ";
                cin >> y.tahunTerbit;
                cout << "Masukkan Konten: ";
                cin >> y.konten;
                Q = alokasiChild(y);
                insertFirst(C, Q);
                cout << "=============================" << endl;
                cout << "Artikel berhasil ditambahkan!" << endl;
                cout << "=============================" << endl;
                break;
            case 3:
                printInfo(L);
                cout << "Masukkan Nama Penulis : ";
                cin >> nama;
                P = findElm_nama(L, nama);
                if (P != NULL) {
                    printInfo(C);
                    cout << "Masukkan Judul Artikel : ";
                    cin >> judul;
                    Q = findElm_judul(C, judul);
                    if (Q != NULL) {
                        if (findElm(R, P, Q) != NULL) {
                            cout << "============================================" << endl;
                            cout << "Relasi antara Penulis dan Artikel sudah ada!" << endl;
                            cout << "============================================" << endl;
                        } else {
                            S = alokasiRelasi(P, Q);
                            insertFirst(R, S);
                            cout << "============================" << endl;
                            cout << "Relasi berhasil ditambahkan!" << endl;
                            cout << "============================" << endl;
                        }
                    } else {
                        cout << "========================" << endl;
                        cout << "Artikel tidak ditemukan." << endl;
                        cout << "========================" << endl;
                    }
                } else {
                cout << "========================" << endl;
                cout << "Penulis tidak ditemukan." << endl;
                cout << "========================" << endl;
                }
                break;

            case 4:
                printInfo(L);
                break;

            case 5:
                printInfo(C);
                break;

            case 6:
                printInfo(R);
                break;

            case 7:
                printInfo(L);
                if (first(L) == NULL) {
                    cout << "===========" << endl;
                    cout << "Data Kosong" << endl;
                    cout << "===========" << endl;
                } else {
                    cout << "Masukkan Nama Penulis : ";
                    cin >> nama;
                    P = findElm_nama(L, nama);
                    if (P != NULL) {
                        deleteRelasiParent(R, P);
                        removePenulis(nama, L, C, R);
                        cout << "===========================================" << endl;
                        cout << "Penulis berhasil dihapus beserta relasinya." << endl;
                        cout << "===========================================" << endl;
                    } else {
                        cout << "========================" << endl;
                        cout << "Penulis tidak ditemukan." << endl;
                        cout << "========================" << endl;
                    }
                }
                break;

            case 8:
                printInfo(C);
                if (first(C) == NULL) {
                    cout << "===========" << endl;
                    cout << "Data Kosong" << endl;
                    cout << "===========" << endl;
                } else {
                    cout << "Masukkan Judul Artikel : ";
                    cin >> judul;
                    Q = findElm_judul(C, judul);
                    if (Q != NULL) {
                        removeArtikel(judul, C, R);
                        cout << "===========================================" << endl;
                        cout << "Artikel berhasil dihapus beserta relasinya." << endl;
                        cout << "===========================================" << endl;
                    } else {
                        cout << "========================" << endl;
                        cout << "Artikel tidak ditemukan." << endl;
                        cout << "========================" << endl;
                    }
                }
                break;

            case 9:
                if (first(R) == NULL) {
                    cout << "=======================================================" << endl;
                    cout << "Tidak ada relasi yang dapat dihapus karena data kosong!" << endl;
                    cout << "=======================================================" << endl;
                } else {
                    string namaParent, judulArtikel;
                    cout << "Masukkan Nama Penulis: ";
                    cin >> namaParent;
                    cout << "Masukkan Judul Artikel: ";
                    cin >> judulArtikel;
                    deleteRelasiParentChild(R, namaParent, judulArtikel);
                }
                break;

            case 10:
                cout << "Masukkan Nama: ";
                cin >> nama;
                searchElm_nama(L, nama);
                break;

            case 11:
                cout << "Masukkan ID: ";
                cin >> nama;
                searchElm_id(L, nama);
                break;

            case 12:
                cout << "Masukkan Email: ";
                cin >> nama;
                searchElm_email(L, nama);
                break;

            case 13:
                cout << "Masukkan Judul Artikel: ";
                cin >> judul;
                searchElm_judul(C, judul);
                break;

            case 14:
                int tahun;
                cout << "Masukkan Tahun Terbit: ";
                cin >> tahun;
                searchElm_tahun(C, tahun);
                break;

            case 15:
                cout << "Masukkan Konten: ";
                cin >> judul;
                searchElm_konten(C, judul);
                break;

            case 16:
                if (first(L) == NULL) {
                    cout << "===========================================================" << endl;
                    cout << "Data penulis kosong. Tidak ada artikel yang dapat dihitung." << endl;
                    cout << "===========================================================" << endl;
                } else {
                    cout << "Masukkan nama penulis yang ingin dihitung artikelnya: ";
                    cin >> nama;
                    address_parent P = findElm_nama(L, nama);
                    if (P != NULL) {
                        countArtikel(R, P);
                    } else {
                        cout << "Penulis dengan nama \"" << nama << "\" tidak ditemukan." << endl;
                    }
                }
                break;

            case 17:
                printInfo(L);
                cout << "Masukkan ID Penulis: ";
                cin >> nama;
                printInfo(C);
                cout << "Masukkan Judul Artikel: ";
                cin >> judul;
                searchArtikelberdasarkanParent(R, nama, judul);
                break;

            case 18:
                cout << "===========================================" << endl;
                cout << "Terima kasih telah menggunakan program ini!" << endl;
                cout << "===========================================" << endl;
                schoice = false;
                break;

            default:
                choice = 18;
                cout << "===================" << endl;
                cout << "Pilihan tidak valid" << endl;
                cout << "===================" << endl;
                break;
        }
    }
}

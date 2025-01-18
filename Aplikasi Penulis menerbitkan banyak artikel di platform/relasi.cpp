#include "relasi.h"
#include <string>
#include <sstream>
void createList(List_relasi &L){
    first(L) = NULL;
}
void insertFirst(List_relasi &L, address_relasi P){
    next(P) = first(L);
    first(L) = P;
}
address_relasi alokasiRelasi( address_parent P, address_child C){
    address_relasi Q = new elmlist_relasi;
    parent(Q) = P;
    child(Q) = C;
    next(Q) = NULL;
    return Q;
}
void deleteFirst(List_relasi &L, address_relasi &P) {
    if (first(L) != NULL) {
        P = first(L);
        first(L) = next(P);
        next(P) = NULL;
        dealokasi(P);
    }
}
void dealokasi(address_relasi &P){
     if (P != NULL) {
        delete P;
        P = NULL;
    }
}
address_relasi findElm(List_relasi L, address_parent P, address_child C) {
    address_relasi Q = first(L);
    while (Q != NULL && (parent(Q) != P || child(Q) != C)) {
        Q = next(Q);
    }
    return Q;
}
void printInfo(List_relasi L) {
    if (first(L) == NULL) {
        cout << "===========" << endl;
        cout << "Data Kosong" << endl;
        cout << "===========" << endl;
        return;
    }

    // Menentukan panjang maksimum untuk setiap kolom (secara manual)
    int maxNama = 4;           // Set default minimal panjang untuk nama
    int maxID = 2;             // Set default minimal panjang untuk ID
    int maxEmail = 5;          // Set default minimal panjang untuk email
    int maxJudulArtikel = 14;  // Set default minimal panjang untuk judul artikel
    int maxTahunTerbit = 12;   // Set default minimal panjang untuk tahun terbit
    int maxKonten = 6;         // Set default minimal panjang untuk konten

    address_relasi P = first(L);

    // Menghitung panjang maksimum untuk kolom-kolom yang ada
    while (P != NULL) {
        if (parent(P) != NULL && child(P) != NULL) {
            // Memperbarui panjang maksimum untuk setiap kolom
            maxNama = max(maxNama, (int)info(parent(P)).nama.length());
            maxID = max(maxID, (int)info(parent(P)).id.length());
            maxEmail = max(maxEmail, (int)info(parent(P)).email.length());
            maxJudulArtikel = max(maxJudulArtikel, (int)info(child(P)).judulArtikel.length());
            maxTahunTerbit = max(maxTahunTerbit, (int)to_string(info(child(P)).tahunTerbit).length());
            maxKonten = max(maxKonten, (int)info(child(P)).konten.length());
        }
        P = next(P);
    }

    // Menampilkan header tabel
    int headerLength = maxNama + maxID + maxEmail + maxJudulArtikel + maxTahunTerbit + maxKonten + 19; // Tambahkan 19 untuk garis pembatas
    string separator(headerLength, '=');  // Membuat garis pembatas sesuai panjang header

    cout << " " << separator << endl;
    cout << "| Nama" << string(maxNama - 4, ' ')
         << " | ID" << string(maxID - 2, ' ')
         << " | Email" << string(maxEmail - 5, ' ')
         << " | Judul Artikel" << string(maxJudulArtikel - 14, ' ')
         << " | Tahun Terbit" << string(maxTahunTerbit - 12, ' ')
         << " | Konten" << string(maxKonten - 6, ' ')
         << " |" << endl;
    cout << " " << separator << endl;

    // Menampilkan data dalam format tabel
    P = first(L);
    while (P != NULL) {
        if (parent(P) != NULL && child(P) != NULL) {
            string nama = info(parent(P)).nama;
            string id = info(parent(P)).id;
            string email = info(parent(P)).email;
            string judulArtikel = info(child(P)).judulArtikel;
            string tahunTerbitStr = to_string(info(child(P)).tahunTerbit);
            string konten = info(child(P)).konten;

            // Menampilkan data dalam tabel dengan penyesuaian panjang kolom
            cout << "| " << nama << string(maxNama - nama.length(), ' ');
            cout << " | " << id << string(maxID - id.length(), ' ');
            cout << " | " << email << string(maxEmail - email.length(), ' ');
            cout << " | " << judulArtikel << string(maxJudulArtikel - judulArtikel.length(), ' ');
            cout << " | " << tahunTerbitStr << string(maxTahunTerbit - tahunTerbitStr.length(), ' ');
            cout << " | " << konten << string(maxKonten - konten.length(), ' ');
            cout << " |" << endl;
        } else {
            cout << "Relasi tidak valid." << endl;
        }
        P = next(P);
    }

    cout << " " << separator << endl;
}
void deleteRelasiParent(List_relasi &L, address_parent P) {
    address_relasi Q = first(L);
    address_relasi temp;

    while (Q != NULL) {
        if (parent(Q) == P) {
            if (Q == first(L)) {
                deleteFirst(L, temp);
                Q = first(L);
            } else {
                address_relasi prevQ = first(L);
                while (next(prevQ) != Q) {
                    prevQ = next(prevQ);
                }
                next(prevQ) = next(Q);
                temp = Q;
                Q = next(Q);
                dealokasi(temp);
            }
        } else {
            Q = next(Q);
        }
    }
}

void deleteRelasiChild(List_relasi &R, address_child P) {
    address_relasi Q = first(R);
    address_relasi prevQ = NULL;

    while (Q != NULL) {
        if (child(Q) == P) {
            if (Q == first(R)) {
                deleteFirst(R, Q);
                Q = first(R);
            } else {
                prevQ = first(R);
                while (next(prevQ) != Q) {
                    prevQ = next(prevQ);
                }
                next(prevQ) = next(Q);
                dealokasi(Q);
                Q = next(prevQ);
            }
        } else {
            prevQ = Q;
            Q = next(Q);
        }
    }
}

void countArtikel(List_relasi R, address_parent P) {
    if (P == NULL) {
        cout << "Penulis tidak ditemukan." << endl;
        return;
    }

    int count = 0;
    address_relasi Q = first(R);

    while (Q != NULL) {
        if (parent(Q) == P) {
            count++;
        }
        Q = next(Q);
    }

    cout << "Jumlah artikel yang terkait dengan penulis " << info(P).nama << " adalah: " << count << endl;
}

void deleteRelasiParentChild(List_relasi &R, string namaParent, string judulArtikel) {
    address_relasi Q = first(R);
    address_relasi prevQ = NULL;
    bool found = false;

    while (Q != NULL) {
        if (info(parent(Q)).nama == namaParent && info(child(Q)).judulArtikel == judulArtikel) {
            found = true;
            if (Q == first(R)) {
                first(R) = next(Q);
            } else {
                prevQ = first(R);
                while (next(prevQ) != Q) {
                    prevQ = next(prevQ);
                }
                next(prevQ) = next(Q);
            }
            dealokasi(Q);
            cout << "Relasi antara \"" << namaParent << "\" dan \"" << judulArtikel << "\" berhasil dihapus!" << endl;
            return;
        }
        Q = next(Q);
    }

    if (!found) {
        cout << "Relasi antara \"" << namaParent << "\" dan \"" << judulArtikel << "\" tidak ditemukan." << endl;
    }
}

void searchArtikelberdasarkanParent(List_relasi &R, string idPenulis, string judulArtikel) {
    bool penulisFound = false, articleFound = false;
    address_relasi relasi = first(R);
    while (relasi != NULL) {
        if (info(parent(relasi)).id == idPenulis) {
            penulisFound = true;
            if (info(child(relasi)).judulArtikel == judulArtikel) {
                articleFound = true;
                cout << "----------------------------------" << endl;
                cout << "Nama Penulis   : " << info(parent(relasi)).nama << endl;
                cout << "ID Penulis     : " << info(parent(relasi)).id << endl;
                cout << "Email Penulis  : " << info(parent(relasi)).email << endl;
                cout << "Judul Artikel  : " << info(child(relasi)).judulArtikel << endl;
                cout << "Tahun Terbit   : " << info(child(relasi)).tahunTerbit << endl;
                cout << "Konten Artikel : " << info(child(relasi)).konten << endl;
                cout << "----------------------------------" << endl;
                break;
            }
        }
        relasi = next(relasi);
    }

    if (!penulisFound) {
        cout << "==================================" << endl;
        cout << "Penulis dengan ID \"" << idPenulis << "\" tidak ditemukan." << endl;
        cout << "==================================" << endl;
    } else if (!articleFound) {
        cout << "==================================" << endl;
        cout << "Artikel dengan judul \"" << judulArtikel << "\" tidak ditemukan pada ID " << idPenulis << endl;
        cout << "==================================" << endl;
    }
}


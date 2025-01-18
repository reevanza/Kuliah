#include "child.h"
#include <string>
#include <sstream>
void createList(List_child &L){
    first(L) = NULL;
    last(L) = NULL;
}
void insertFirst(List_child &L, address_child P) {
    address_child Q = findElm_judul(L, info(P).judulArtikel);
    if (Q != NULL) {
        cout << "=================================================" << endl;
        cout << "Judul Artikel sudah ada, tidak dapat menambahkan." << endl;
        cout << "=================================================" << endl;
        return;
    }
    if (first(L) == NULL) {
        first(L) = P;
        last(L) = P;
    } else {
        next(P) = first(L);
        prev(first(L)) = P;
        first(L) = P;
    }
}

void deleteFirst(List_child &L, address_child &P){
    P = first(L);
    if (next(P) == NULL) {
        first(L) = NULL;
        last(L) = NULL;
    } else {
        first(L) = next(P);
        prev(first(L)) = NULL;
    }
    next(P) = NULL;
    prev(P) = NULL;
}
address_child alokasiChild(infotype_child x){
    address_child P = new elmlist_child;
    info(P) = x;
    next(P) = NULL;
    prev(P) = NULL;
    return P;
}
void deleteAfter (List_child &L, address_child Prec, address_child &P){
    P = next(Prec);
    if (P == last(L)) {
        last(L) = Prec;
        next(Prec) = NULL;
    } else {
        next(Prec) = next(P);
        prev(next(P)) = Prec;
    }
    next(P) = NULL;
    prev(P) = NULL;
}
void deleteLast(List_child &L, address_child &P){
    P = last(L);
    if (prev(P) == NULL) {
        first(L) = NULL;
        last(L) = NULL;
    } else {
        last(L) = prev(P);
        next(last(L)) = NULL;
    }
    prev(P) = NULL;
    next(P) = NULL;
}
void removeArtikel(string judulArtikel, List_child &L, List_relasi &R){
    address_child P = first(L);
    while (P != NULL && info(P).judulArtikel != judulArtikel) {
        P = next(P);
    }
    if (P != NULL) {
        deleteRelasiChild(R, P);
        if (P == first(L)) {
            deleteFirst(L, P);
        } else if (P == last(L)) {
            deleteLast(L, P);
        } else {
            address_child Q = prev(P);
            deleteAfter(L, Q, P);
        }
        dealokasi(P);
        cout << "=========================" << endl;
        cout << "Artikel berhasil dihapus." << endl;
        cout << "=========================" << endl;
    } else {
        cout << "========================" << endl;
        cout << "Artikel tidak ditemukan." << endl;
        cout << "========================" << endl;
    }
}

void dealokasi(address_child &P) {
    if (P != NULL) {
        delete P;
        P = NULL;
    }
}

void searchElm_judul(List_child L, string x) {
    address_child P = first(L);
    bool found = false;
    cout << "==================================" << endl;
    cout << "Pencarian berdasarkan judul artikel: " << x << endl;
    cout << "==================================" << endl;

    while (P != NULL) {
        if (info(P).judulArtikel == x) {
            found = true;
            cout << "----------------------------------" << endl;
            cout << "Judul Artikel : " << info(P).judulArtikel << endl;
            cout << "Tahun Terbit  : " << info(P).tahunTerbit << endl;
            cout << "Konten        : " << info(P).konten << endl;
            cout << "Lokasi memory : " << P << endl;
            cout << "----------------------------------" << endl;
        }
        P = next(P);
    }

    if (!found) {
        cout << "====================" << endl;
        cout << "Data tidak ditemukan" << endl;
        cout << "====================" << endl;
    }
}


void searchElm_tahun(List_child L, int x) {
    address_child P = first(L);
    bool found = false;
    cout << "==================================" << endl;
    cout << "Pencarian berdasarkan tahun terbit: " << x << endl;
    cout << "==================================" << endl;

    while (P != NULL) {
        if (info(P).tahunTerbit == x) {
            found = true;
            cout << "----------------------------------" << endl;
            cout << "Judul Artikel : " << info(P).judulArtikel << endl;
            cout << "Tahun Terbit  : " << info(P).tahunTerbit << endl;
            cout << "Konten        : " << info(P).konten << endl;
            cout << "Lokasi memory : " << P << endl;
            cout << "----------------------------------" << endl;
        }
        P = next(P);
    }

    if (!found) {
        cout << "====================" << endl;
        cout << "Data tidak ditemukan" << endl;
        cout << "====================" << endl;
    }
}


void searchElm_konten(List_child L, string x) {
    address_child P = first(L);
    bool found = false;

    cout << "==================================" << endl;
    cout << "Pencarian berdasarkan konten: " << x << endl;
    cout << "==================================" << endl;

    while (P != NULL) {
        if (info(P).konten == x) {
            found = true;
            cout << "----------------------------------" << endl;
            cout << "Judul Artikel : " << info(P).judulArtikel << endl;
            cout << "Tahun Terbit  : " << info(P).tahunTerbit << endl;
            cout << "Konten        : " << info(P).konten << endl;
            cout << "Lokasi memory : " << P << endl;
            cout << "----------------------------------" << endl;
        }
        P = next(P);
    }

    if (!found) {
        cout << "====================" << endl;
        cout << "Data tidak ditemukan" << endl;
        cout << "====================" << endl;
    }
}



address_child findElm_judul(List_child L, string x){
    address_child P = first(L);
    while(P != NULL && info(P).judulArtikel != x){
        P = next(P);
    }
    return P;
}

void printInfo(List_child L) {
    if (first(L) == NULL) {
        cout << "===========" << endl;
        cout << "Data Kosong" << endl;
        cout << "===========" << endl;
        return;
    }
    // Panjang header
    int minJudulArtikel = 14, minTahunTerbit = 12, minKonten = 6;
    int maxJudulArtikel = minJudulArtikel, maxTahunTerbit = minTahunTerbit, maxKonten = minKonten;

    // Menghitung panjang maksimum untuk setiap kolom berdasarkan data
    address_child P = first(L);
    while (P != NULL) {
        maxJudulArtikel = max(maxJudulArtikel, (int)info(P).judulArtikel.length());
        maxTahunTerbit = max(maxTahunTerbit, (int)to_string(info(P).tahunTerbit).length());
        maxKonten = max(maxKonten, (int)info(P).konten.length());
        P = next(P);
    }

    // Membuat garis pembatas
    int headerLength = maxJudulArtikel + maxTahunTerbit + maxKonten + 8;
    string separator(headerLength, '=');

    // Menampilkan header tabel
    cout << " " << separator << endl;
    cout << "| Judul Artikel" << string(maxJudulArtikel - minJudulArtikel, ' ')
         << " | Tahun Terbit" << string(maxTahunTerbit - minTahunTerbit, ' ')
         << " | Konten" << string(maxKonten - minKonten, ' ') << " |" << endl;
    cout << " " << separator << endl;

    // Menampilkan data dalam tabel
    P = first(L);
    while (P != NULL) {
        string judulArtikel = info(P).judulArtikel;
        string tahunTerbitStr = to_string(info(P).tahunTerbit);
        string konten = info(P).konten;
        cout << "| " << judulArtikel << string(maxJudulArtikel - judulArtikel.length(), ' ')
             << "| " << tahunTerbitStr << string(maxTahunTerbit - tahunTerbitStr.length(), ' ')
             << " | " << konten << string(maxKonten - konten.length(), ' ')
             << " |" << endl;
        P = next(P);
    }
    cout << " " << separator << endl;

}

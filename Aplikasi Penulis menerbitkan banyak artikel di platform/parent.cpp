#include "parent.h"
#include "relasi.h"
#include <string>
#include <sstream>
void createList(List_parent &L){
    first(L) = NULL;
}
void insertFirst(List_parent &L, address_parent P) {
    if (findElm_id(L, info(P).id) != NULL) {
        cout << "======================================" << endl;
        cout << "ID sudah ada, tidak dapat menambahkan." << endl;
        cout << "======================================" << endl;
        dealokasi(P);
        return;
    }
    next(P) = first(L);
    first(L) = P;
    cout << "=====================================" << endl;
    cout << "Penulis berhasil ditambahkan di awal!" << endl;
    cout << "=====================================" << endl;
}
void insertAfter(List_parent &L, address_parent Prec, address_parent P) {
    if (findElm_id(L, info(P).id) != NULL) {
        cout << "======================================" << endl;
        cout << "ID sudah ada, tidak dapat menambahkan." << endl;
        cout << "======================================" << endl;
        dealokasi(P);
        return;
    }
    next(P) = next(Prec);
    next(Prec) = P;
}

void insertLast(List_parent &L, address_parent P) {
    if (findElm_id(L, info(P).id) != NULL) {
        cout << "======================================" << endl;
        cout << "ID sudah ada, tidak dapat menambahkan." << endl;
        cout << "======================================" << endl;
        dealokasi(P);
        return;
    }
    if (first(L) == NULL) {
        first(L) = P;
    } else {
        address_parent Q = first(L);
        while (next(Q) != NULL) {
            Q = next(Q);
        }
        next(Q) = P;
    }
    cout << "======================================" << endl;
    cout << "Penulis berhasil ditambahkan di akhir!" << endl;
    cout << "======================================" << endl;
}

address_parent findElm_nama(List_parent L, string x){
    address_parent P = first(L);
    while(P != NULL && info(P).nama != x){
        P = next(P);
    }
    return P;
}

address_parent findElm_id(List_parent L, string x) {
    address_parent P = first(L);
    while (P != NULL && info(P).id != x) {
        P = next(P);
    }
    return P;
}

void deleteFirst(List_parent &L, address_parent &P){
    P = first(L);
    first(L) = next(P);
    next(P) = NULL;
}
void deleteLast(List_parent &L, address_parent &P){
    address_parent Q = first(L);
    while(next(next(Q)) != NULL){
        Q = next(Q);
    }
    P = next(Q);
    next(Q) = NULL;
}
void deleteAfter(List_parent &L, address_parent Prec, address_parent &P){
    P = next(Prec);
    next(Prec) = next(P);
    next(P) = NULL;
}
void removePenulis(string nama, List_parent &L, List_child &C, List_relasi &R) {
    address_parent P = findElm_nama(L, nama);
    if (P != NULL) {
        deleteRelasiParent(R, P);
        address_parent temp;
        if (P == first(L)) {
            deleteFirst(L, temp);
        } else {
            address_parent prec = first(L);
            while (next(prec) != P) {
                prec = next(prec);
            }
            deleteAfter(L, prec, temp);
        }
        dealokasi(temp);
        cout << "=========================" << endl;
        cout << "Penulis berhasil dihapus." << endl;
        cout << "=========================" << endl;
    } else {
        cout << "========================" << endl;
        cout << "Penulis tidak ditemukan." << endl;
        cout << "========================" << endl;
    }
}

address_parent alokasiParent(infotype_parent x){
    address_parent P = new elmlist_parent;
    info(P) = x;
    next(P) = NULL;
    return P;
}
void dealokasi(address_parent &P){
    delete P;
}

void searchElm_nama(List_parent L, string x) {
    address_parent P = first(L);
    bool found = false;

    cout << "==================================" << endl;
    cout << "Pencarian berdasarkan nama: " << x << endl;
    cout << "==================================" << endl;

    while (P != NULL) {
        if (info(P).nama == x) {
            found = true;
            cout << "----------------------------------" << endl;
            cout << "Nama          : " << info(P).nama << endl;
            cout << "ID            : " << info(P).id << endl;
            cout << "Email         : " << info(P).email << endl;
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


void searchElm_id(List_parent L, string x) {
    address_parent P = first(L);
    bool found = false;

    cout << "==================================" << endl;
    cout << "Pencarian berdasarkan ID: " << x << endl;
    cout << "==================================" << endl;

    while (P != NULL) {
        if (info(P).id == x) {
            found = true;
            cout << "----------------------------------" << endl;
            cout << "Nama          : " << info(P).nama << endl;
            cout << "ID            : " << info(P).id << endl;
            cout << "Email         : " << info(P).email << endl;
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


void searchElm_email(List_parent L, string x) {
    address_parent P = first(L);
    bool found = false;

    cout << "==================================" << endl;
    cout << "Pencarian berdasarkan email: " << x << endl;
    cout << "==================================" << endl;

    while (P != NULL) {
        if (info(P).email == x) {
            found = true;
            cout << "----------------------------------" << endl;
            cout << "Nama          : " << info(P).nama << endl;
            cout << "ID            : " << info(P).id << endl;
            cout << "Email         : " << info(P).email << endl;
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

void printInfo(List_parent L) {
    if (first(L) == NULL) {
        cout << "===========" << endl;
        cout << "Data Kosong" << endl;
        cout << "===========" << endl;
        return;
    }

    // Menentukan panjang maksimum untuk setiap kolom (secara manual)
    int maxNama = 4;    // Set default minimal panjang untuk nama
    int maxID = 2;      // Set default minimal panjang untuk ID
    int maxEmail = 5;   // Set default minimal panjang untuk email

    // Menghitung panjang maksimum untuk nama, ID, dan email secara manual
    address_parent P = first(L);
    while (P != NULL) {
        // Memperbarui panjang maksimum nama
        if (info(P).nama.length() > maxNama) {
            maxNama = info(P).nama.length();
        }
        // Memperbarui panjang maksimum ID
        if (info(P).id.length() > maxID) {
            maxID = info(P).id.length();
        }
        // Memperbarui panjang maksimum email
        if (info(P).email.length() > maxEmail) {
            maxEmail = info(P).email.length();
        }
        P = next(P);
    }

    // Menampilkan header tabel
    int headerLength = maxNama + maxID + maxEmail + 8; // Total panjang header berdasarkan kolom
    string separator(headerLength, '=');  // Membuat garis pembatas sesuai panjang header

    cout << " " << separator << endl;
    cout << "| Nama" << string(maxNama - 4, ' ') << " | ID" << string(maxID - 2, ' ') << " | Email" << string(maxEmail - 5, ' ') << " |" << endl;
    cout <<  " " << separator << endl;

    // Menampilkan data dalam format tabel
    P = first(L);
    while (P != NULL) {
        string nama = info(P).nama;
        string id = info(P).id;
        string email = info(P).email;

        // Menampilkan data dalam tabel dengan penyesuaian panjang kolom
        cout << "| " << nama;
        for (int i = nama.length(); i < maxNama; i++) cout << " "; // Menyesuaikan lebar kolom nama

        cout << " | " << id;
        for (int i = id.length(); i < maxID; i++) cout << " "; // Menyesuaikan lebar kolom id

        cout << " | " << email;
        for (int i = email.length(); i < maxEmail; i++) cout << " "; // Menyesuaikan lebar kolom email

        cout << " |" << endl;
        P = next(P);
    }

    cout << " " << separator << endl;
}

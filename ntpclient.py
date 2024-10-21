import socket
import time

def main():
    # Mengatur alamat IP dan port NTP server
    server_address = ('localhost', 123)  # Ganti 'localhost' dengan alamat server jika perlu

    # Membuat socket UDP
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    try:
        # Mengirimkan permintaan waktu ke server
        print("Mengirimkan permintaan waktu ke server...")
        client_socket.sendto(b"", server_address)

        # Menerima waktu dari server
        data, _ = client_socket.recvfrom(1024)

        # Mengubah waktu yang diterima menjadi format int
        server_time = int(data.decode('utf-8'))

        # Menampilkan waktu server dan waktu lokal
        print(f"Waktu dari NTP server: {time.ctime(server_time)}")
        print(f"Waktu lokal: {time.ctime(time.time())}")

    except Exception as e:
        print("Terjadi kesalahan:", e)

    finally:
        client_socket.close()

if __name__ == "__main__":
    main()

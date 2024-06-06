<!DOCTYPE html>
<html>
<head>
    <title>List Kategori Artikels</title>
</head>
<body>
    <h1>Daftar Kategori Artikel</h1>

    <!-- Tambah tombol untuk menambah kategori artikel -->
    <a href="{{ route('kategori-artikel.create') }}" class="btn btn-primary">Tambah Kategori Artikel</a>

    <!-- Tampilkan daftar kategori artikel -->
    <ul>
        @foreach($kategoriArtikels as $kategoriArtikel)
            <li>{{ $kategoriArtikel['nama'] }}</li>
            <!-- Tambah tombol untuk mengedit dan menghapus kategori artikel -->
            <a href="{{ route('kategori-artikel.edit', $kategoriArtikel['id']) }}" class="btn btn-primary">Edit</a>
            <form action="{{ route('kategori-artikel.destroy', $kategoriArtikel['id']) }}" method="POST">
                @csrf
                @method('DELETE')
                <button type="submit" class="btn btn-danger">Hapus</button>
            </form>
        @endforeach
    </ul>
</body>
</html>

<!DOCTYPE html>
<html>
<head>
    <title>Ubah Kategori Artikel</title>
</head>
<body>
<h1>Edit Kategori Artikel</h1>
    
    <form action="{{ route('kategori-artikel.update', $kategoriArtikel['id']) }}" method="POST" enctype="multipart/form-data">
        @csrf
        @method('PUT')
        <div class="form-group">
            <label for="nama">Nama:</label>
            <input type="text" name="nama" class="form-control" value="{{ $kategoriArtikel['nama'] }}" required>
        </div>
        <button type="submit" class="btn btn-primary">Update</button>
    </form>
</body>
</html>

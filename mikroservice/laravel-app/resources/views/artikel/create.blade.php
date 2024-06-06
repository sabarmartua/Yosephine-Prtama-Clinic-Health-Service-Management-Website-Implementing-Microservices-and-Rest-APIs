<!DOCTYPE html>
<html>

<head>
    <title>Tambah Artikel Baru</title>
    <!-- Link CSS Bootstrap -->
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
    <style>
        body {
            background-color: #f8f9fa;
            margin-top: 50px;
        }
        .container {
            margin-top: 20px;
        }
        .card {
            border-radius: 10px;
            box-shadow: 0px 2px 10px rgba(0, 0, 0, 0.1);
        }
        .card-header {
            background-color: #007bff;
            color: #fff;
            border-bottom: none;
            border-top-left-radius: 10px;
            border-top-right-radius: 10px;
        }
        .form-group {
            margin-bottom: 20px;
        }
        label {
            font-weight: bold;
        }
        input[type="text"],
        textarea,
        select {
            border: 1px solid #ced4da;
            border-radius: 5px;
            padding: 10px;
            width: 100%;
            transition: border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
        }
        input[type="text"]:focus,
        textarea:focus,
        select:focus {
            border-color: #007bff;
            outline: 0;
            box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
        }
        .btn-primary {
            background-color: #007bff;
            border-color: #007bff;
            padding: 10px 30px;
            font-size: 16px;
            border-radius: 5px;
        }
        .btn-primary:hover {
            background-color: #0056b3;
            border-color: #0056b3;
        }
    </style>
</head>

<body>
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Tambah Artikel Baru</div>

                    <div class="card-body">
                        <form method="POST" action="{{ route('artikel.store') }}" enctype="multipart/form-data">
                            @csrf
                            <div class="form-group">
                                <label for="nama">Nama Artikel</label>
                                <input type="text" class="form-control" id="nama" name="nama" required>
                            </div>

                            <div class="form-group">
                                <label for="konten">Konten</label>
                                <textarea class="form-control" id="konten" name="konten" rows="6" required></textarea>
                            </div>

                            <div class="form-group">
                                <label for="kategori_id">Kategori</label>
                                <select class="form-control" id="kategori_id" name="kategori_id" required>
                                    <option value="">Pilih Kategori</option>
                                    @if(!empty($categories))
                                    @foreach ($categories as $category)
                                    <option value="{{ $category['id'] }}">{{ $category['nama'] }}</option>
                                    @endforeach
                                    @else
                                    <option value="" disabled>Servis Kategori Sedang Tidak Tersedia</option>
                                    @endif
                                </select>
                            </div>
                            <div class="form-group">
                                <label for="gambar">Gambar</label>
                                <input type="file" class="form-control-file" id="gambar" name="gambar" required>
                            </div>

                            <button type="submit" class="btn btn-primary">Tambah Artikel</button>
                        </form>

                    </div>
                </div>
            </div>
        </div>
    </div>
</body>

</html>

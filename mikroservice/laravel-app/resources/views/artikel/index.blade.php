<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <!-- Tambahkan link CSS Bootstrap di sini jika belum termasuk -->
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
    <style>
        /* Tambahkan gaya tambahan di sini jika diperlukan */
        .card {
            margin-bottom: 20px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }

        .card-body {
            padding: 20px;
        }
    </style>
</head>
<body>
<div class="container">
    @if(session('status'))
        <div class="alert alert-success">{{ session('status') }}</div>
    @endif

    @if(session('error'))
        <div class="alert alert-danger">{{ session('error') }}</div>
    @endif

    <h1>Daftar Artikel</h1>

    <div class="row">
        @foreach($artikels as $artikel)
            <div class="col-md-4">
                <div class="card">
                    <div class="card-body">
                        <h5 class="card-title">{{ $artikel['nama'] }}</h5>
                        <p class="card-text">{{ $artikel['konten'] }}</p>
                        <p class="card-text"><strong>Kategori:</strong> {{ $artikel['kategori'] }}</p>
                        <img src="{{ asset('uploads/images/' . $artikel['gambar']) }}" class="card-img-top" alt="{{ $artikel['nama'] }}">
                        <div class="mt-3">
                            <a href="{{ route('artikel.edit', $artikel['id']) }}" class="btn btn-primary">Edit</a>
                            <form action="{{ route('artikel.destroy', $artikel['id']) }}" method="POST" style="display: inline-block;">
                                @csrf
                                @method('DELETE')
                                <button type="submit" class="btn btn-danger" onclick="return confirm('Are you sure you want to delete this artikel?')">Delete</button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        @endforeach
    </div>
</div>
</body>
</html>

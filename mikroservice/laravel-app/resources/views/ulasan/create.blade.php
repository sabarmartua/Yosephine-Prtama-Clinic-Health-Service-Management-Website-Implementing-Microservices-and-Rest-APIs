<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tambah Ulasan Baru</title>
</head>
<body>
    <h1>Tambah Ulasan Baru</h1>

    @if($errors->any())
        <div style="color: red;">
            <ul>
                @foreach($errors->all() as $error)
                    <li>{{ $error }}</li>
                @endforeach
            </ul>
        </div>
    @endif

    <form action="{{ route('ulasan.store') }}" method="post">
        @csrf
        <div>
            <label for="isiUlasan">Isi Ulasan:</label><br>
            <textarea id="isiUlasan" name="isiUlasan" rows="4" cols="50"></textarea>
        </div>
        <br>
        <button type="submit">Tambah Ulasan</button>
    </form>
</body>
</html>

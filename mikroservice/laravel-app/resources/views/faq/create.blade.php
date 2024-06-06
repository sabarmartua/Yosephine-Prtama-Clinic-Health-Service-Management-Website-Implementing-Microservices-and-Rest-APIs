<!DOCTYPE html>
<html>
<head>
    <title>Tambah FAQ</title>
</head>
<body>
    <h1>Tambah FAQ</h1>

    @if ($errors->any())
        <div style="color: red">
            <ul>
                @foreach ($errors->all() as $error)
                    <li>{{ $error }}</li>
                @endforeach
            </ul>
        </div>
    @endif

    <form action="{{ route('faq.store') }}" method="POST">
        @csrf
        <div>
            <label for="pertanyaan">Pertanyaan:</label>
            <input type="text" id="pertanyaan" name="pertanyaan">
        </div>
        <div>
            <label for="jawaban">Jawaban:</label>
            <input type="text" id="jawaban" name="jawaban">
        </div>
        <button type="submit">Tambah FAQ</button>
    </form>
</body>
</html>

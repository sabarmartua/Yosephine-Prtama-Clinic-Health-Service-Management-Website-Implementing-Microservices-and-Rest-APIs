<!DOCTYPE html>
<html>
<head>
    <title>Edit FAQ</title>
</head>
<body>
    <h1>Edit FAQ</h1>

    @if ($errors->any())
        <div style="color: red">
            <ul>
                @foreach ($errors->all() as $error)
                    <li>{{ $error }}</li>
                @endforeach
            </ul>
        </div>
    @endif

    <form action="{{ route('faq.update', $faq['id']) }}" method="POST">
        @csrf
        @method('PUT')
        <div>
            <label for="pertanyaan">Pertanyaan:</label>
            <input type="text" id="pertanyaan" name="pertanyaan" value="{{ $faq['pertanyaan'] }}">
        </div>
        <div>
            <label for="jawaban">Jawaban:</label>
            <input type="text" id="jawaban" name="jawaban" value="{{ $faq['jawaban'] }}">
        </div>
        <button type="submit">Simpan Perubahan</button>
    </form>
</body>
</html>

<!DOCTYPE html>
<html>
<head>
    <title>Daftar FAQ</title>
</head>
<body>
    <h1>Daftar FAQ</h1>

    @if(session('success'))
        <div style="color: green">{{ session('success') }}</div>
    @endif

    @if(session('error'))
        <div style="color: red">{{ session('error') }}</div>
    @endif

    <a href="{{ route('faq.create') }}">Tambah FAQ</a>

    <ul>
        @foreach($faqs as $faq)
            <li>
                <strong>Pertanyaan:</strong> {{ $faq['pertanyaan'] }} - 
                <strong>Jawaban:</strong> {{ $faq['jawaban'] }}
                <a href="{{ route('faq.edit', $faq['id']) }}">Edit</a>
                <form action="{{ route('faq.destroy', $faq['id']) }}" method="POST">
                    @csrf
                    @method('DELETE')
                    <button type="submit">Hapus</button>
                </form>
            </li>
        @endforeach
    </ul>
</body>
</html>

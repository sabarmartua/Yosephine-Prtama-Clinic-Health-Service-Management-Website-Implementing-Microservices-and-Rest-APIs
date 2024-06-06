<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Artikel;
use App\Models\Kategori;
use GuzzleHttp\Client;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Validator;

class ArtikelController extends Controller
{
    protected $apiUrl;
    private $messages;
    private $client;

    public function __construct()
    {
        // Tentukan URL API Go CRUD
        $this->apiUrl = 'http://192.168.154.117:8084/api/artikel'; // Ganti dengan alamat IP perangkat yang sesuai // Ganti dengan URL API Go Anda
        $this->messages = [
            'required' => 'The :attribute field is required.',
            'image' => 'The :attribute must be an image.',
            'mimes' => 'The :attribute must be a file of type: :values.',
            'max' => 'The :attribute may not be greater than :max kilobytes.',
        ];

        $this->client = new Client(); // Initialize GuzzleHttp\Client
    }

    public function index()
    {
        try {
            $response = Http::get($this->apiUrl . '/all');
            $artikelsData = $response->json();

            if (!is_array($artikelsData)) {
                return redirect()->back()->with('error', 'Error fetching data');
            }

            // Ambil data kategori dari API
            $kategoriResponse = Http::get("http://192.168.154.117:8083/api/kategori-artikel");
            $categories = $kategoriResponse->json();

            // Jika data kategori tidak berupa array, tangani kesalahan
            if (!is_array($categories)) {
                return redirect()->back()->with('error', 'Error fetching data');
            }

            // Ubah format data artikel dengan menambahkan informasi kategori
            $artikels = [];
            foreach ($artikelsData as $artikel) {
                $kategori_id = $artikel['kategori_id'];
                $kategori = collect($categories)->firstWhere('id', $kategori_id);
                $artikel['kategori'] = $kategori ? $kategori['nama'] : 'Unknown';
                $artikels[] = $artikel;
            }

            return view('artikel.index', compact('artikels'));
        } catch (\Exception $e) {
            return view('error')->with('message', 'Server is currently down. Please try again later.');
        }
    }

    public function create()
    {
        try {
            $categoryResponse = $this->client->get("http://192.168.154.117:8083/api/kategori-artikel");
            $categories = json_decode($categoryResponse->getBody(), true);
    
            // Periksa apakah $categories adalah array yang valid
            if (!is_array($categories)) {
                // Jika tidak, berikan array kosong agar tidak terjadi kesalahan
                $categories = [];
            }
    
            return view('artikel.create', compact('categories'));
        } catch (\Exception $e) {
            // Tangani kesalahan dengan mengembalikan pesan error
            return view('artikel.create')->with('error', 'Error fetching categories');
        }
    }
    

    public function store(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'nama' => 'required|string|min:3|max:255',
            'konten' => 'required|string',
            'kategori_id' => 'required|numeric',
            'gambar' => 'required|image|mimes:jpeg,png,jpg|max:20480',
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }

        $gambar = $request->file('gambar');
        $gambarName = time() . '_' . $gambar->getClientOriginalName();
        $gambar->move(public_path('uploads/images'), $gambarName);

        $parameter = [
            'nama' => $request->nama,
            'konten' => $request->konten,
            'kategori_id' => (int)$request->kategori_id,
            'gambar' => $gambarName
        ];

        try {
            $client = new Client();
            $url = "http://192.168.154.117:8084/api/artikel/create";
            $response = $client->request('POST', $url, [
                'headers' => ['Content-Type' => 'application/json'],
                'body' => json_encode($parameter)
            ]);
            $content = $response->getBody()->getContents();
            $contentArray = json_decode($content, true);

            if (isset($contentArray['status']) && $contentArray['status'] == false) {
                return redirect()->back()->with('error', $contentArray['message']);
            }

            $data = $contentArray['data'];

            return redirect()->route('artikel.index')->with('message', 'Tambah Artikel Berhasil');
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage())->withInput();
        }
    }

    public function edit($id)
    {
        try {
            $response = $this->client->get($this->apiUrl . '/' . $id);
            $artikel = json_decode($response->getBody(), true);

            $categoryResponse = $this->client->get("http://192.168.154.117:8083/api/kategori-artikel");
            $categories = json_decode($categoryResponse->getBody(), true);

            if (!is_array($categories) || !is_array($artikel)) {
                return redirect()->back()->with('error', 'Error fetching data');
            }

            return view('artikel.edit', compact('artikel', 'categories'));
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage());
        }
    }

    public function update(Request $request, $id)
    {
        $validator = Validator::make($request->all(), [
            'nama' => 'required|string|min:3|max:255',
            'konten' => 'required|string',
            'kategori_id' => 'required|numeric',
            'gambar' => 'nullable|image|mimes:jpeg,png,jpg|max:20480',
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }

        $parameter = [
            'nama' => $request->nama,
            'konten' => $request->konten,
            'kategori_id' => (int)$request->kategori_id,
        ];

        // Handle image upload if provided
        if ($request->hasFile('gambar')) {
            $gambar = $request->file('gambar');
            $gambarName = time() . '_' . $gambar->getClientOriginalName();
            $gambar->move(public_path('uploads/images'), $gambarName);
            $parameter['gambar'] = $gambarName;
        }

        try {
            $client = new Client();
            $url = "http://192.168.154.117:8084/api/artikel/update/{$id}";
            $response = $client->request('PUT', $url, [
                'headers' => ['Content-Type' => 'application/json'],
                'body' => json_encode($parameter)
            ]);
            $content = $response->getBody()->getContents();
            $contentArray = json_decode($content, true);

            if (isset($contentArray['status']) && $contentArray['status'] == false) {
                return redirect()->back()->with('error', $contentArray['message']);
            }

            return redirect()->route('artikel.index')->with('message', 'Update Artikel Berhasil');
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage())->withInput();
        }
    }

    public function destroy($id)
    {
        try {
            $response = $this->client->delete($this->apiUrl . '/delete/' . $id);

            if ($response->getStatusCode() == 204) {
                return redirect()->route('artikel.index')->with('status', 'Artikel deleted successfully');
            } else {
                return redirect()->back()->with('error', 'Error deleting artikel');
            }
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage());
        }
    }
}

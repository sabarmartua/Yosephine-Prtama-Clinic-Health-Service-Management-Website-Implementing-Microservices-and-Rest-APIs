<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use GuzzleHttp\Client;
use Illuminate\Support\Facades\Validator;

class KategoriArtikelController extends Controller
{
    protected $apiUrl;
    private $client;

    public function __construct()
    {
        // Tentukan URL API Go CRUD
        $this->apiUrl = 'http://192.168.154.117:8083/api/kategori-artikel'; // Ganti dengan URL API Go Anda

        // Initialize the Guzzle client
        $this->client = new Client();
    }

    public function index()
    {
        try {
            $response = $this->client->get($this->apiUrl . '/');

            if ($response->getStatusCode() === 200) {
                $kategoriArtikels = json_decode($response->getBody(), true);
                return view('kategori-artikel.index', compact('kategoriArtikels'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return view('error')->with('message', 'Server is currently down. Please try again later.');
        }
    }

    public function create()
    {
        return view('kategori-artikel.create');
    }

    public function store(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'nama' => 'required',
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }

        try {
            $response = $this->client->post($this->apiUrl . '/', [
                'json' => [
                    'nama' => $request->nama,
                ]
            ]);

            return redirect()->route('kategori-artikel.index')->with('success', 'Kategori artikel berhasil ditambahkan.');
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage())->withInput();
        }
    }

    public function edit($id)
    {
        try {
            $response = $this->client->get($this->apiUrl . '/' . $id);
            $kategoriArtikel = json_decode($response->getBody(), true);

            // Pastikan $kategoriArtikel merupakan array yang memiliki properti 'id'
            if (!isset($kategoriArtikel['id'])) {
                throw new \Exception('Invalid response data');
            }

            return view('kategori-artikel.edit', compact('kategoriArtikel'));
        } catch (\Exception $e) {
            return redirect()->route('kategori-artikel.index')->with('error', $e->getMessage());
        }
    }


    public function update(Request $request, $id)
    {
        $validator = Validator::make($request->all(), [
            'nama' => 'required',
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors($validator)->withInput();
        }

        try {
            $response = $this->client->put($this->apiUrl . '/' . $id, [
                'json' => [
                    'nama' => $request->nama,
                ]
            ]);

            return redirect()->route('kategori-artikel.index')->with('success', 'Kategori artikel berhasil diperbarui.');
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage())->withInput();
        }
    }

    public function destroy($id)
    {
        try {
            $response = $this->client->delete($this->apiUrl . '/' . $id);

            return redirect()->route('kategori-artikel.index')->with('success', 'Kategori artikel berhasil dihapus.');
        } catch (\Exception $e) {
            return redirect()->route('kategori-artikel.index')->with('error', $e->getMessage());
        }
    }
}

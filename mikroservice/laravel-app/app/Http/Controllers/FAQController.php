<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class FAQController extends Controller
{
    protected $apiUrl;

    public function __construct()
    {
        $this->apiUrl = 'http://192.168.154.117:8086/api/faq'; // Sesuaikan dengan URL servis FAQ Anda
    }

    public function index()
    {
        try {
            $response = Http::get($this->apiUrl . '/all');

            if ($response->status() == 200) {
                $faqs = $response->json();
                return view('faq.index', compact('faqs'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return view('error')->with('message', 'Server is currently down. Please try again later.');
        }
    }

    public function create()
    {
        return view('faq.create');
    }

    public function store(Request $request)
    {
        $validatedData = $request->validate([
            'pertanyaan' => 'required|string',
            'jawaban' => 'required|string',
        ]);

        try {
            $response = Http::post($this->apiUrl . '/create', $validatedData);

            if ($response->status() == 201) {
                return redirect()->route('faq.index')->with('success', 'FAQ berhasil ditambahkan.');
            } else {
                throw new \Exception('Failed to add data to API');
            }
        } catch (\Exception $e) {
            return redirect()->route('faq.index')->with('error', $e->getMessage())->withInput();
        }
    }

    public function edit($id)
    {
        try {
            $response = Http::get($this->apiUrl . '/' . $id);

            if ($response->status() == 200) {
                $faq = $response->json();
                return view('faq.edit', compact('faq'));
            } else {
                throw new \Exception('Failed to fetch data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('faq.index')->with('error', $e->getMessage());
        }
    }

    public function update(Request $request, $id)
    {
        $validatedData = $request->validate([
            'pertanyaan' => 'required|string',
            'jawaban' => 'required|string',
        ]);

        try {
            $response = Http::put($this->apiUrl . '/update/' . $id, $validatedData);

            if ($response->status() == 200) {
                return redirect()->route('faq.index')->with('success', 'FAQ berhasil diperbarui.');
            } else {
                throw new \Exception('Failed to update data to API');
            }
        } catch (\Exception $e) {
            return redirect()->back()->with('error', $e->getMessage())->withInput();
        }
    }

    public function destroy($id)
    {
        try {
            $response = Http::delete($this->apiUrl . '/delete/' . $id);

            if ($response->status() == 200) {
                return redirect()->route('faq.index')->with('success', 'FAQ berhasil dihapus.');
            } else {
                throw new \Exception('Failed to delete data from API');
            }
        } catch (\Exception $e) {
            return redirect()->route('faq.index')->with('error', $e->getMessage());
        }
    }
}

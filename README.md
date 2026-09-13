# Anti-Virtualization via WMI Hardware Fingerprinting

<img width="1672" height="941" alt="Banner Image" src="https://github.com/user-attachments/assets/218a2d68-ce31-4ba3-b9bf-263803f799cb" />

A lightweight demonstration showing how software can inspect underlying hardware artifacts to detect virtual machines and automated security sandboxes.

Rather than relying on invasive hooks or suspicious low-level API calls, this implementation leverages native **Windows Management Instrumentation (WMI)** queries to verify whether the system is running on bare metal or inside a hypervisor. It evaluates key hardware descriptors against known vendor signatures:

* **System Information:** Manufacturer and Model (`Win32_ComputerSystem`)
* **Firmware Data:** BIOS Manufacturer and Version (`Win32_BIOS`)
* **Motherboard Details:** BaseBoard Manufacturer and Product (`Win32_BaseBoard`)

If typical virtualization strings (such as VMware, VirtualBox, QEMU, or Xen) are detected, the environment is flagged as an analysis sandbox.

---

### 📖 Accompanying Write-Up

For a comprehensive breakdown of the evasion mechanics, theory, and sandbox behavior, check out the full article:

**[Read the Full Article: Anti-Virtualization Explained: Detecting Hypervisors Through Native WMI Queries](https://www.google.com/search?q=INSERT_YOUR_ARTICLE_LINK_HERE)**
